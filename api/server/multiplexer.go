package server

import (
	"log"
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

type job struct {
	jobID   string
	request []byte
	reply   chan []byte
}

type wire interface {
	loadZMQ(msg [][]byte) error
	sendZMQ(socket *zmq.Socket) (total int, err error)
}

/*
 * The partitionRequest is the message sent from this (the session manager)
 * over ZMQ to the scheduler/job partitioner, which decides what data to
 * retrieve from storage (manifest-server in current vocabulary)
 */
type partitionRequest struct {
	// Return-address that flows through to make sure that data is returned to
	// the correct node that manages the session
	address string
	jobID string
	request []byte
}

/*
 * The partialResult is this model internal representation of the message
 * sent by the worker nodes when a task is done.
 *
 * The payload being an opaque blob of bytes is very useful for testing, since
 * the ZMQ and channel messaging infrastructure now does not depend on the
 * payload, and instead of structured protobuf messages we can send strings to
 * compare.
 */
type partialResult struct {
	jobID string
	payload []byte
}

/*
 * The routedPartialRequest is a more faithful representation of what the
 * worker nodes *actually* send - they need to include a return address to the
 * node that holds the session (this program, really). However, ZMQ at some
 * point strips this address as a part of its routing protocol, and the rest of
 * the application sees the message as partialResult.
 */
type routedPartialResult struct {
	address string
	partial partialResult
}

/*
 * The make/send functions are stupid helpers to help formalise the protocol
 * for communication with other parts of oneseismic, and provide a canonical
 * way of formatting messages for both the wire (over ZMQ) and over go channels
 */
func newPartitionRequest(j *job, address string) *partitionRequest {
	return &partitionRequest {
		address: address,
		jobID: j.jobID,
		request: j.request,
	}
}

func (p *partitionRequest) sendZMQ(socket *zmq.Socket) (total int, err error) {
	return socket.SendMessage(p.address, p.jobID, p.request)
}

/*
 * Parse a partition request as it is delivered in a ZMQ multipart message.
 * While this currently has no error checking, or really does anything
 * sophisticated, it's the canonical way to obtain a partitionRequest from a
 * multipart message, and *the* go reference for what the messages from the
 * fragment/worker looks like.
 */
func (p *partitionRequest) loadZMQ(msg [][]byte) error {
	if len(msg) != 3 {
		return fmt.Errorf("len(msg) = %d; want 3", len(msg))
	}

	p.address = string(msg[0])
	p.jobID = string(msg[1])
	p.request = msg[2]
	return nil
}

func (p *partialResult) loadZMQ(msg [][]byte) error {
	if len(msg) != 2 {
		return fmt.Errorf("len(msg) = %d; want 2", len(msg))
	}
	p.jobID = string(msg[0])
	p.payload = msg[1]
	return nil
}

func (p *partialResult) sendZMQ(socket *zmq.Socket) (total int, err error) {
	return socket.SendMessage(p.jobID, p.payload)
}

func (p *routedPartialResult) loadZMQ(msg [][]byte) error {
	if len(msg) != 3 {
		return fmt.Errorf("len(msg) = %d; want 3", len(msg))
	}
	p.address = string(msg[0])
	err := p.partial.loadZMQ(msg[1:])
	if err != nil {
		return fmt.Errorf("routedPartialResult.partial: %s", err.Error())
	}
	return nil
}

func (p *routedPartialResult) sendZMQ(socket *zmq.Socket) (total int, err error) {
	return socket.SendMessage(p.address, p.partial.jobID, p.partial.payload)
}

func multiplexer(jobs chan job, address string, reqNdpt string, repNdpt string) {
	req, err := zmq.NewSocket(zmq.PUSH)

	if err != nil {
		log.Fatal(err)
	}

	req.Bind(reqNdpt)

	rep := make(chan partialResult)
	go func() {
		r, err := zmq.NewSocket(zmq.DEALER)

		if err != nil {
			log.Fatal(err)
		}

		r.SetIdentity(address)
		r.Bind(repNdpt)

		var partial partialResult
		for {
			m, err := r.RecvMessageBytes(0)

			if err != nil {
				log.Fatal(err)
			}

			partial.loadZMQ(m)
			rep <- partial
		}
	}()

	// TODO: Clean up in case a reply never arrives?
	replyChnls := make(map[string]chan []byte)

	for {
		select {
		case r := <-rep:
			rc := replyChnls[r.jobID]
			rc <- r.payload
			delete(replyChnls, r.jobID)

		case j := <-jobs:
			part := newPartitionRequest(&j, address)
			replyChnls[j.jobID] = j.reply
			part.sendZMQ(req)
		}
	}
}
