# TM5GProject
Simulator of the 5G Core Network (specifically, the registration procedure).

## Features

### Registration

The simulator is comprised of seven Network Functions and contains all the services and models required by the registration procedure. It  was built primarily with NF interaction (and the data that is part of these interactions) in mind. Therefore, it is possible to alter every data field that is sent at each step of the procedure. This allows us to analyse the characteristics of different types of data flowing through the system.

### UE Context Generation

As part of the registration procedure, information about the User Equipment is randomly generated based on the required formats of the data. This information is used as the basis for the rest of the procedure. This approach was chosen because it allows for more granular control of the data used in the simulation than attaching a UE-RAN simulator.

### Logging

Important information about NF requests and responses (such as size, body contents, duration, metadata) is recorded and saved to log files. The purpose of this is to use the information as part of an analysis. 

### Cloud Database Storage

MongoDB is used for the storage needs of this simulator. During the registration procedure, data is inserted, updated, retrieved, and deleted as necessary.

## Installation

```
go get github.com/tommcclymont/TM5GProject
```

## Usage

After navigating to the project directory, run the following command to start the NF servers:

```go
go run AMF/amf.go
go run AUSF/ausf.go
go run EIR/eir.go
go run OldAMF/oldamf.go
go run PCF/pcf.go
go run SMF/smf.go
go run UDM/udm.go
```

Now that all the servers are running, begin the registration procedure:

```go
go run registration.go
```
