# NiGoAPI

## Purpose

The aim of this project is to provide a convenient Go client to interact with NiFi components ([NiFi](https://nifi.apache.org/) Api and [NiFi Registry](https://nifi.apache.org/registry.html) API)

## Documentations

- [NiFi client](pkg/registry/README.md)
- [NiFi Registry client](pkg/registry/README.md)

## Generate swagger file

To generate the swagger files for a specific <version\>, run the following commands:

```console
git checkout tags/rel/nifi-<version> -b rel/nifi-<version>
```

Build the `swagger.json` for **NiFi**:

```console
cd nifi-framework-bundle/nifi-framework/nifi-web/nifi-web-api
mvn clean install -T2C -DskipTests
```

Now available at `nifi-framework-bundle/nifi-framework/nifi-web/nifi-web-api/target/nifi-web-api-2.5.0/docs/rest-api/swagger.json`.

Build the `swagger.json` for **NiFi Registry**:

```console
cd nifi-registry/nifi-registry-core/nifi-registry-web-api
mvn clean install -T2C -DskipTests
```

Now available at `nifi-registry/nifi-registry-core/nifi-registry-web-api/target/nifi-registry-web-api-2.5.0/docs/rest-api/swagger.json`.

## Generate client

To generate the clients please run following command: 

```console
./regenerate.sh
```

Manual changes: 

- In [`model_version_control_information_dto`](./pkg/nifi/model_version_control_information_dto.go), replace:

  ```golang
  // The version of the flow
	Version string `json:"version,omitempty"`
  ```

  by:

  ```golang
  // The version of the flow
	Version interface{} `json:"version,omitempty"`
  ```
  
- In [`model_revision_dto.go`](./pkg/nifi/model_revision_dto.go), replace: 

  ```golang
  Version int64 `json:"version,omitempty"`
  ```
  
  by:
  
  ```golang
  Version *int64 `json:"version,omitempty"`
  ```  

- In [`model_label_dto.go`](./pkg/nifi/model_label_dto.go), replace: 
  
  ```golang
	// The text that appears in the label.
	Label string `json:"label,omitempty"`
  ```
    
  by:
  
  ```golang
	// The text that appears in the label.
	Label string `json:"label"`
  ```

- In [`LabelDto.md`](./pkg/nifi/docs/LabelDto.md), replace: 
  
  ```markdown
  **Label** | **string** | The text that appears in the label. | [optional] [default to null]
  ```
  
  by:
  
  ```markdown
  **Label** | **string** | The text that appears in the label. | [required] [default to null]
  ```
  
- In [`model_parameter_dto.go`](./pkg/nifi/model_parameter_dto.go), replace: 
  
  ```golang
	// The description of the Parameter
	Description string `json:"description,omitempty"`
	// The value of the Parameter
	Value string `json:"value,omitempty"`
  ```
    
  by:
  
  ```golang
	// The description of the Parameter
	Description *string `json:"description,omitempty"`
	// The value of the Parameter
	Value *string `json:"value"`
  ```

- In [`ParameterDto.md`](./pkg/nifi/docs/ParameterDto.md), replace: 
  
  ```markdown
  **Value** | **string** | The value of the Parameter | [optional] [default to null]
  ```
  
  by:
  
  ```markdown
  **Value** | **string** | The value of the Parameter | [required] [default to null]
  ```

- In [`model_parameter_context_dto`](./pkg/nifi/model_parameter_context_dto.go), replace: 
  
  ```golang
	// The Parameters for the Parameter Context
  Parameters []ParameterEntity `json:"parameters,omitempty"`
  // The Process Groups that are bound to this Parameter Context
  BoundProcessGroups []ProcessGroupEntity `json:"boundProcessGroups,omitempty"`
  // A list of references of Parameter Contexts from which this one inherits parameters
  InheritedParameterContexts []ParameterContextReferenceEntity `json:"inheritedParameterContexts,omitempty"`
  ```
  
  by:
  
  ```golang
	// The Parameters for the Parameter Context
  Parameters []ParameterEntity `json:"parameters"`
  // The Process Groups that are bound to this Parameter Context
  BoundProcessGroups []ProcessGroupEntity `json:"boundProcessGroups"`
  // A list of references of Parameter Contexts from which this one inherits parameters
  InheritedParameterContexts []ParameterContextReferenceEntity `json:"inheritedParameterContexts"`
  ```

- In [`ParameterContextDto`](./pkg/nifi/docs/ParameterContextDto.md), replace: 
  
  ```markdown
  **Parameters** | [**[]ParameterEntity**](ParameterEntity.md) | The Parameters for the Parameter Context | [optional] [default to null]
  **BoundProcessGroups** | [**[]ProcessGroupEntity**](ProcessGroupEntity.md) | The Process Groups that are bound to this Parameter Context | [optional] [default to null]
  **InheritedParameterContexts** | [**[]ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) | A list of references of Parameter Contexts from which this one inherits parameters | [optional] [default to null]
  ```
  
  by:
  
  ```markdown
  **Parameters** | [**[]ParameterEntity**](ParameterEntity.md) | The Parameters for the Parameter Context | [required] [default to null]
  **BoundProcessGroups** | [**[]ProcessGroupEntity**](ProcessGroupEntity.md) | The Process Groups that are bound to this Parameter Context | [required] [default to null]
  **InheritedParameterContexts** | [**[]ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) | A list of references of Parameter Contexts from which this one inherits parameters | [required] [default to null]
  ```

- In [`model_position_dto`](./pkg/nifi/model_position_dto.go), replace: 
  
  ```golang
  // The x coordinate.
  X float64 `json:"x,omitempty"`
  // The y coordinate.
  Y float64 `json:"y,omitempty"`
  ```
    
  by:
    
  ```golang
  // The x coordinate.
  X float64 `json:"x"`
  // The y coordinate.
  Y float64 `json:"y"`
  ```

- In [`PositionDto`](./pkg/nifi/docs/PositionDto.md), replace: 
  
  ```markdown
  **X** | **float64** | The x coordinate. | [optional] [default to null]
  **Y** | **float64** | The y coordinate. | [optional] [default to null]
  ```
    
  by:
  
  ```markdown
  **X** | **float64** | The x coordinate. | [required] [default to null]
  **Y** | **float64** | The y coordinate. | [required] [default to null]
  ```
  
- replace 

  ```golang
  // The timestamp of when the request was submitted
  SubmissionTime time.Time `json:"submissionTime,omitempty"`
  // The timestamp of when the request was last updated
  LastUpdated time.Time `json:"lastUpdated,omitempty"`
  ```
 
  by: 
  
  ```golang
  // The timestamp of when the request was submitted
  SubmissionTime string `json:"submissionTime,omitempty"`
  // The timestamp of when the request was last updated
  LastUpdated string `json:"lastUpdated,omitempty"`
  ```

- replace

  ```golang
	// The index of the bend point where to place the connection label.
  LabelIndex int32 `json:"labelIndex,omitempty"`
  ```

  by:

  ```golang
	// The index of the bend point where to place the connection label.
  LabelIndex int32 `json:"labelIndex"`
  ```

## Credits

This client is: 

- inspired by the python client [nipyapi](https://github.com/Chaffelson/nipyapi)
- using [swagger-codegen](https://swagger.io/tools/swagger-codegen/)