package oneseismic;

import java.io.InputStream;
import java.io.IOException;
import java.nio.ByteBuffer;
import java.nio.FloatBuffer;
import java.nio.ByteOrder;
import java.util.Arrays;

public class SimpleDecoder {

    private static int BUFFER_SIZE = 10240;

    private int[] shape;
    private float[] data_buffer;
    
    
    public SimpleDecoder() {
        
    }

    public float[] getData() {
        return data_buffer;
    }

    public int[] getShape() {
        return shape;
    }

    public boolean process(InputStream input) {
        boolean result = false;
        
        DecoderJNI decoder = new DecoderJNI();
        byte[] buffer = new byte[BUFFER_SIZE];

        String[] attrs = null;
        try {
            //First read until have header
            for (int read_size = input.read(buffer);
                 read_size != -1;
                 read_size = input.read(buffer)) {
                decoder.process(buffer,read_size);
                attrs = decoder.getAttr();
                if (attrs != null) {
                    this.shape = decoder.getShape();
                    this.shape = Arrays.copyOfRange(this.shape, 1, this.shape.length);    
                    break;
                }
            }
            
            if ((attrs == null) || (this.shape == null)){
                return false;
            }
            
            int total_size=1;
            for (int i = 0; i < this.shape.length; i++) {
                total_size=total_size *this.shape[i];
            }
            
            ByteBuffer shared_buffer = ByteBuffer.allocateDirect(total_size*4);
            decoder.setWriter("data",shared_buffer);

            for (int read_size = input.read(buffer);
                 read_size != -1;
                 read_size = input.read(buffer)) {
                decoder.process(buffer,read_size);
            }

        
            FloatBuffer float_buffer =  shared_buffer.order(ByteOrder.LITTLE_ENDIAN).asFloatBuffer();
            
            this.data_buffer = new float[total_size];
            for(int i=0;i<data_buffer.length;i++) {
                data_buffer[i] = float_buffer.get(i);
            } 
            result = true;
        }
        catch (IOException ioe) {
            System.err.println("Problem reading stream");
        }
        return result;
    }

    
}
