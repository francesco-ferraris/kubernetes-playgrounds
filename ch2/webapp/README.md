# Description
Deploy the application built in ch1 playground using an Helm Chart
## Usage
```makefile
make install-helm
```
Install Helm
<br/>
<br/>
```makefile
make lint
```
Lint the Helm Chart
<br/>
<br/>
```makefile
make template
```
Render the chart with the default values in the `output` directory
<br/>
<br/>
```makefile
make deploy
```
Deploy the Chart using the default values
<br/>
<br/>
## Cleanup
```makefile
make delete
```
Delete the Chart
