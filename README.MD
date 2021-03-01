# NiGoAPI

## Purpose

The aim of this project is to provide a convenient Go client to interact with NiFi components ([NiFi](https://nifi.apache.org/) Api and [NiFi Registry](https://nifi.apache.org/registry.html) API)

## Documentations

- [NiFi client](pkg/registry/README.md)
- [NiFi Registry client](pkg/registry/README.md)

## Generate client

To generate the clients please run following command : 

```console
./regenerate.sh
```

Manual changes : 

- In function `UploadTemplate()` of `api_process_groups.go` replace :

  ```golang
  localVarPostBody = &localVarOptionals.Body.Value()
  ```
  
  by :
  
  ```golang
  localVarPostBody = localVarOptionals.Body.Value()
  ```  
  
- In `model_revision_dto`, replace : 

  ```golang
  Version int64 `json:"version,omitempty"`
  ```
  
  by :
  
  ```golang
  Version *int64 `json:"version,omitempty"`
  ```  
  
  - In `model_parameter_dto`, replace : 
  
    ```golang
    Value string `json:"value,omitempty"`
    ```
    
    by :
    
    ```golang
    Value *string `json:"value,omitempty"`
    ``` 
  
- replace 

  ```golang
  // The timestamp of when the request was submitted
  SubmissionTime Time.time `json:"submissionTime,omitempty"`
  // The timestamp of when the request was last updated
  LastUpdated Time.time `json:"lastUpdated,omitempty"`
  ```
 
  by : 
  
  ```golang
  // The timestamp of when the request was submitted
  SubmissionTime string `json:"submissionTime,omitempty"`
  // The timestamp of when the request was last updated
  LastUpdated string `json:"lastUpdated,omitempty"`
  ```
  
 
 

## Credits

This client is : 

- inspired by the python client [nipyapi](https://github.com/Chaffelson/nipyapi)
- using [swagger-codegen](https://swagger.io/tools/swagger-codegen/)