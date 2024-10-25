# Real Image Challenge 2016

## How to use:

The linux executable is present in bin folder. 

To build code from source use:
```
go build -o ./bin/app main.go
```

To start the server run the below command:
```
./bin/app
```

You can directly run the code using command below:
```
go run main.go
```

HTTP server will start at `http://localhost:8080`

API documentation is hosted at `http://localhost:8080/docs/index.html` 

You can use the Swagger UI to test the APIs as well. Alternatively you can import the Postman collection in Postman applictaion stored in docs folders to use the APIs.

## Update Swagger

This implementation uses github.com/swaggo/swag/cmd/swag to generate swagger documentation. You can install it by running below command.

```
go install github.com/swaggo/swag/cmd/swag@latest
```  

To update swagger documentation, just update the comments on the API handlers and run the below command

```
swag init
```
---
### Format for region string

Whenever you are passing region string e.g. in Permissions or checking if region is valid for a given ditributor follow the format in same order as below.

`CityCode-ProvinceCode-CountryCode`

e.g. YELUR-KA-IN, TN-IN, US, etc.

---

## Prolem Statement
In the cinema business, a feature film is usually provided to a regional distributor based on a contract for exhibition in a particular geographical territory.

Each authorization is specified by a combination of included and excluded regions. For example, a distributor might be authorzied in the following manner:
```
Permissions for DISTRIBUTOR1
INCLUDE: INDIA
INCLUDE: UNITEDSTATES
EXCLUDE: KARNATAKA-INDIA
EXCLUDE: CHENNAI-TAMILNADU-INDIA
```
This allows `DISTRIBUTOR1` to distribute in any city inside the United States and India, *except* cities in the state of Karnataka (in India) and the city of Chennai (in Tamil Nadu, India).

At this point, asking your program if `DISTRIBUTOR1` has permission to distribute in `CHICAGO-ILLINOIS-UNITEDSTATES` should get `YES` as the answer, and asking if distribution can happen in `CHENNAI-TAMILNADU-INDIA` should of course be `NO`. Asking if distribution is possible in `BANGALORE-KARNATAKA-INDIA` should also be `NO`, because the whole state of Karnataka has been excluded.

Sometimes, a distributor might split the work of distribution amount smaller sub-distiributors inside their authorized geographies. For instance, `DISTRIBUTOR1` might assign the following permissions to `DISTRIBUTOR2`:

```
Permissions for DISTRIBUTOR2 < DISTRIBUTOR1
INCLUDE: INDIA
EXCLUDE: TAMILNADU-INDIA
```
Now, `DISTRIBUTOR2` can distribute the movie anywhere in `INDIA`, except inside `TAMILNADU-INDIA` and `KARNATAKA-INDIA` - `DISTRIBUTOR2`'s permissions are always a subset of `DISTRIBUTOR1`'s permissions. It's impossible/invalid for `DISTRIBUTOR2` to have `INCLUDE: CHINA`, for example, because `DISTRIBUTOR1` isn't authorized to do that in the first place. 

If `DISTRIBUTOR2` authorizes `DISTRIBUTOR3` to handle just the city of Hubli, Karnataka, India, for example:
```
Permissions for DISTRIBUTOR3 < DISTRIBUTOR2 < DISTRIBUTOR1
INCLUDE: HUBLI-KARNATAKA-INDIA
```
Again, `DISTRIBUTOR2` cannot authorize `DISTRIBUTOR3` with a region that they themselves do not have access to. 

We've provided a CSV with the list of all countries, states and cities in the world that we know of - please use the data mentioned there for this program. *The codes you see there may be different from what you see here, so please always use the codes in the CSV*. This Readme is only an example. 

Write a program in any language you want (If you're here from Gophercon, use Go :D) that does this. Feel free to make your own input and output format / command line tool / GUI / Webservice / whatever you want. Feel free to hold the dataset in whatever structure you want, but try not to use external databases - as far as possible stick to your langauage without bringing in MySQL/Postgres/MongoDB/Redis/Etc.

To submit a solution, fork this repo and send a Pull Request on Github. 

For any questions or clarifications, raise an issue on this repo and we'll answer your questions as fast as we can.


