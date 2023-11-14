package banking

import (
	"github.com/manmohansharma21/bankpoc/banking-lib/logger"
	"github.com/manmohansharma21/bankpoc/banking/app"
)

func main() {
	//log.Println("Starting our application...")
	logger.Info("Starting our application...") //Equivalent to fmt.Println in standard library
	app.Start()
}

// To run the application: SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWD=root DB_ADDR=localhost DB_PORT=3306 DB_NAME=banking go run main.go ---> For passing the env vars; or we can simple EXPORT indivisually or
// or via shell script. That shell will not be commited to git, as it will be almost same as baking conf creds with application code.

// DTO ------creates-----> Domain Object
/*
Hexagonal Architecture==Dependency Inversion ===> Highly maintainable code, agnostic to the outsider world.

[User]-----Login Request----->[Auth Server]-----Token Response--->[User]----Request resource with Token in the Header----> [Resource Server (Banking)]----------Verify Token and Request----->[Auth Server]-----Token verification Request----->[Resource Server banking]-----Response----->[User]

As REST APIs are stateless means they should not maintain the state; every API call coming to the server must be self contained means conrainng information for authorization for its data, verified every time. Fast and efficient mechanism is JWT token and their verification = append(As REST APIs are stateless means they should not maintain the state; every API call coming to the server must be self contained means conrainng information for authorization for its data, verified every time. Fast and efficient mechanism is JWT token and their verification.
JWT==JSON WEB TOKEN------> [Header.Payload.Signature]
Header contains type and hashing algorithm (HS256 etc). Payload contains information we want to transmit like issuer, expiry etc. These informations known as claims. Registerd claims are which reserved for token iformation like expiry, issued at etc. We can add our private claims also. Once token parsed, it will contain all info.
Signature is formed by Hash of the header and payload and the secret. Secret is the key stored in the server, the token is signed by this key, maing token secure. This signature ensures unedited claims as updating claims will not have same sercret, hence placing signature properly in server is the security. Should not contain sensitive info in payload claims as that is modifiable = append(Signature is formed by Hash of the header and payload and the secret. Secret is the key stored in the server, the token is signed by this key, maing token secure. This signature ensures unedited claims as updating claims will not have same sercret, hence placing signature properly in server is the security. Should not contain sensitive info in payload claims as that is modifiable.

Authentication:- Verifying identity
Authorization:- granting or denying access to specific resources

*/
/*
* Receiver receives the dependencies, that is why its name, to provide dependency injection we attach receivers rather than pssing value via arguments.
* ORM simplifies the DB interactions by providing drivers which have direct functions using GO structures rather than specifying the SQL queries.
* Layers to Manintain During Development:
1. Application Layer
2. DTO Layer
3. Storage Layer
4. Authentication Layer
5. Service OR Infrastructure Layer
*
*/

/*
Cross Cutting Concerns: Logging, Error handling and Security. These are the integral parts of the
applications across all layers, need to be
designed with minimal or low depenedencies, hence gloablly accessible to all parts of our application.
So, need to get log files or errors in suitable specific structure.

With growing penetration of cloud native architecture, there are even more sources of collection of log data,
where we need to know origin of the logs out of several container or virtual machines.
Logging collector platforms such as(Sumologic or loggly paper trail or any service built over elastic search)
collect these logs. Having a structure helps to eliminate multiple parser need, and can directly be sent to log aggregator platforms = append(collect these logs. Having a structure helps to eliminate multiple parser need, and can directly be sent
to log aggregator platforms being JSON commonly used format.
uber/zap or logger used for structured logging. user/zap known for its performance.
*/

/*
DTO layers helps in modularity and helps in prevening domain objects scattering all over the different layers = append(DTO layers helps in modularity and helps in prevening domain objects scattering all over the different layers = append(DTO layers helps in modularity and helps in prevening domain objects scattering all over the different layers = append(DTO layers helps in modularity and helps in prevening domain objects scattering all over the different layers.
As our domain should not be exposed to the outsider world, DTO layer will help here. Domain object and DTO hold different responsibilities where domain object is at service side layer and DTO is at user side layer.
DTO is for service layer(server side) object for data transformation whereas domain object is the user side object.
Domain has complete knowledge for constructing DTO.
*/
