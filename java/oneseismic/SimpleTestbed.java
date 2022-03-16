package oneseismic;


import java.util.HashMap;
import java.io.IOException;
import java.io.InputStream;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpHeaders;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.net.http.HttpRequest.BodyPublishers;
import java.time.Duration;
import java.nio.ByteBuffer;
import java.nio.FloatBuffer;
import java.nio.ByteOrder;
    
import org.json.simple.JSONArray;
import org.json.simple.JSONObject;
import org.json.simple.parser.JSONParser;
import org.json.simple.parser.ParseException;

public class SimpleTestbed {


    public static String sas = "sv=2020-08-04&ss=b&srt=sco&sp=rwdlacix&se=2022-03-29T14:37:23Z&st=2022-03-15T07:37:23Z&spr=https&sig=kEP0DkkHN95xutDsCD1KBMl%2BTjv7jpks2m1AEfcJQLY%3D";

   public static String lineno_query = """
                {
                    cube(id: \"0e1753a6cbc402e464311782a96b02bf3ac4a2e5\")
                        {
                            linenumbers
                                }
                }""";

    public static String slice_query = """
{
    cube(id: \"0e1753a6cbc402e464311782a96b02bf3ac4a2e5\")
    {
        sliceByLineno(dim: 1, lineno: 24362)
    }
}
    """;

    

 

    private static final HttpClient httpClient = HttpClient.newBuilder()
        .version(HttpClient.Version.HTTP_1_1)
        .connectTimeout(Duration.ofSeconds(60))
        .build();

    public static HttpRequest request(String url,String sas,String query)
    {
        String full_url = url+"?"+sas;
        JSONObject obj = new JSONObject();
        obj.put("query",query);

        String query_json = obj.toJSONString();
        
        HttpRequest request = HttpRequest.newBuilder()
            .POST(BodyPublishers.ofString(query_json))
            .uri(URI.create(full_url))
            .build();
        return request;
    }
    
    public static void main(String[] args)
    {


        try {
            //HttpRequest request = request("http://localhost/os/graphql",sas,lineno_query);
            HttpRequest request = request("http://localhost/os/graphql",sas,slice_query);
            HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());

            JSONParser parser = new JSONParser();
            
            JSONObject jsonObject = (JSONObject) parser.parse(response.body());
            JSONObject cube =(JSONObject) ((JSONObject)jsonObject.get("data")).get("cube");
            HashMap <String,String > results = (HashMap<String,String>) cube.get("sliceByLineno");
            System.out.println(results.get("url"));
            String base_result_url = (String) results.get("url");
            String token = (String) results.get("key");
            String full_request_url = "http://localhost/os/"+base_result_url+"/stream";
            
            HttpRequest cube_request = HttpRequest.newBuilder()
                .GET()
                .setHeader("Authorization","Bearer "+token)
                .uri(URI.create(full_request_url))
                .build();
            HttpResponse<InputStream> sresponse = httpClient.send(cube_request, HttpResponse.BodyHandlers.ofInputStream());
            InputStream data_stream = sresponse.body();

            SimpleDecoder simple_dec = new SimpleDecoder();
            if (simple_dec.process(data_stream))  {
                
                float[] result = simple_dec.getData();
                int[] shape = simple_dec.getShape();
                int row_width =shape[2];
                for (int i=result.length-row_width;i<result.length;i++) {
                    System.out.println(i+" "+result[i]);
                }
            } else {
                System.err.println("Unable to process result");
            }
            
        }
        catch (Exception e)
            {
                e.printStackTrace();
            }
    
        
        
        
        
        System.out.println("Fin");
    }


}
