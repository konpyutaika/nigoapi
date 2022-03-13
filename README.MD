# NiGoAPI

## Purpose

The aim of this project is to provide a convenient Go client to interact with NiFi components ([NiFi](https://nifi.apache.org/) Api and [NiFi Registry](https://nifi.apache.org/registry.html) API)

## Documentations

- [NiFi client](pkg/registry/README.md)
- [NiFi Registry client](pkg/registry/README.md)

## Generate Swagger

Build NiFi source and fetch the `swagger.json` in `nifi-nar-bundles/nifi-framework-bundle/nifi-framework/nifi-web/nifi-web-api/target/swagger-ui/swagger.json`.

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
    // The value of the Parameter
    Value string `json:"value,omitempty"`
    // The description of the Parameter
	  Description string `json:"description,omitempty"`
    ```
    
    by :
    
    ```golang
    // The value of the Parameter
    Value *string `json:"value"`
    // The description of the Parameter
	  Description *string `json:"description,omitempty"`
    ```

- In `model_parameter_context_dto`, replace : 
  
    ```golang
    Parameters []ParameterEntity `json:"parameters,omitempty"`
    // The Process Groups that are bound to this Parameter Context
    BoundProcessGroups []ProcessGroupEntity `json:"boundProcessGroups,omitempty"`
    // A list of references of Parameter Contexts from which this one inherits parameters
    InheritedParameterContexts []ParameterContextReferenceEntity `json:"inheritedParameterContexts,omitempty"`
    ```
    
    by :
    
    ```golang
    Parameters []ParameterEntity `json:"parameters"`
    // The Process Groups that are bound to this Parameter Context
    BoundProcessGroups []ProcessGroupEntity `json:"boundProcessGroups"`
    // A list of references of Parameter Contexts from which this one inherits parameters
    InheritedParameterContexts []ParameterContextReferenceEntity `json:"inheritedParameterContexts"`
    ```

- In `ParameterContextDto`, replace : 
  
    ```markdown
    **Parameters** | [**[]ParameterEntity**](ParameterEntity.md) | The Parameters for the Parameter Context | [optional] [default to null]
    **InheritedParameterContexts** | [**[]ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) | A list of references of Parameter Contexts from which this one inherits parameters | [optional] [default to null]
    ```
    
    by :
    
    ```markdown
    **Parameters** | [**[]ParameterEntity**](ParameterEntity.md) | The Parameters for the Parameter Context | [required] [default to null]
    **InheritedParameterContexts** | [**[]ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) | A list of references of Parameter Contexts from which this one inherits parameters | [required] [default to null]
    ```


- In `model_position_dto`, replace : 
  
    ```golang
    // The x coordinate.
    X float64 `json:"x,omitempty"`
    // The y coordinate.
    Y float64 `json:"y,omitempty"`
    ```
    
    by :
    
    ```golang
    // The x coordinate.
    X float64 `json:"x"`
    // The y coordinate.
    Y float64 `json:"y"`
    ```

- In `Position`, replace : 
  
    ```markdown
    **X** | **float64** | The x coordinate. | [optional] [default to null]
    **Y** | **float64** | The y coordinate. | [optional] [default to null]
    ```
    
    by :
    
    ```markdown
    **X** | **float64** | The x coordinate. | [required] [default to null]
    **Y** | **float64** | The y coordinate. | [required] [default to null]
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