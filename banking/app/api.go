package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manmohansharma21/bankpoc/banking/domain"
	"github.com/manmohansharma21/bankpoc/banking/service"
)

type APIServer struct {
	addr string // The server address.
	port string
}

// Every server must have a Run()
func (s *APIServer) Run() { // Dependency injection via receiver argument
	fmt.Println("Server2 started on port....", s.port)

	router := mux.NewRouter()

	// wiring everything together
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient) // Using same connection pool and same hanndle in all repositories.

	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub(dbClient))} // Injected this stub implementation, later on, will inject real database adapter at the time of wiring our application.

	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}
	// Adding Routes to Endpoints
	//router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet) //Method Matcher using constant defined in net/http library
	router.HandleFunc("/customers", ch.getAllCustomersJSON).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet) // URL segmenting, retrieving information URL.
	// Add segment to it. Request matcher by adding regular expression which will only take numeric id and onwards, else 404 case
	// Segment names are used to create a map of route variables, and can be retrieved using mux.Vars(r) where we need to parse the request inside it,
	// so this function returns map of all the segment names, i.e., vars:=mux.Vars(r); fmt.Fprint(w, vars["customer_id"])

	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAcount).Methods(http.MethodPost) // URL segmenting, retrieving information URL.

	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost) //Same endpoint used for Post request

	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).
		Methods(http.MethodPost)

	fmt.Printf("Server starting on address %s:%s", s.addr, s.port)

	//Starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", s.addr, s.port), router)) //or http.ListenAndServe("localhost:8000", nil); To Start the Server, spinning up the server.
	// SERVER_ADDRESS=localhost SERVER_PORT=8282 go run main.go
}

// newAPIServer: Helping function
func newCustomerAPIServer(addr string, port string) *APIServer {
	return &APIServer{
		addr: addr,
		port: port,
	}
}

/*
1.) When deploying a Go application on Amazon Elastic Kubernetes Service (EKS), the `http.ListenAndServe` function and its parameters are handled differently compared to traditional server setups. In an EKS deployment, your Go application is typically containerized and deployed within a Kubernetes cluster. The networking and service exposure are managed by Kubernetes. Here's how it works:

1. **Containerization**: You package your Go application in a Docker container. You define a Docker image that contains your Go application and its dependencies.

2. **Kubernetes Service**: In Kubernetes, you define a Service resource to expose your application to the network. The Service resource is used to load balance traffic to your application pods. You typically create a LoadBalancer service if you want to expose your application to the internet. Here's an example of a Service definition in a Kubernetes YAML file:

    ```yaml
    apiVersion: v1
    kind: Service
    metadata:
      name: my-app-service
    spec:
      selector:
        app: my-app
      ports:
        - protocol: TCP
          port: 80
          targetPort: 8080
      type: LoadBalancer
    ```

    - `selector` specifies which pods should be associated with this service.
    - `ports` define the ports to open on the service, where `port` is the port exposed externally, and `targetPort` is the port on your application container.
    - `type` is set to `LoadBalancer` to request a cloud provider's load balancer (e.g., AWS Elastic Load Balancer) to distribute incoming traffic.

3. **Kubernetes Deployment**: You define a Deployment or StatefulSet to manage the deployment and scaling of your application pods. These resources specify how many replicas (pod instances) of your application should run.

4. **Port Configuration**: Within your Go application, you can use the standard port, such as `":8080"`, for the HTTP server to listen on. You don't need to specify an IP address.

5. **Kubernetes Networking**: Kubernetes handles networking, routing, and load balancing. You don't directly specify the IP address or port like you would in a traditional server setup. Your application communicates with other services and pods using DNS and service names.

6. **Environment Variables**: You can use environment variables to configure your Go application. For example, you can set environment variables for your application's configuration, such as database connection strings or feature flags. These environment variables can be set within the pod spec in your Kubernetes deployment.

Here's a simplified Go application example:

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, EKS!")
	})

	http.ListenAndServe(":8080", nil)
}
```

In summary, when deploying a Go application on EKS, you don't need to specify an IP address and port in the `http.ListenAndServe` function as you would in a traditional server setup. The IP address and port are managed by Kubernetes and the Service resource. Your application can listen on a standard port (e.g., `":8080"`) and communicate with other services using service names. The service definition and networking configuration in Kubernetes take care of routing and load balancing for your application.

Yes, you can choose any available port instead of port 8080 in your Go application when deploying it on Kubernetes or any other environment. However, there are a few important considerations to keep in mind:

1. **Port Availability**: Ensure that the port you choose is available and not already in use by other processes on the same server or container. Common ports, such as 80 and 443 for HTTP and HTTPS, are often used by other services, so it's a good practice to choose a less commonly used port.

2. **Privileged Ports**: Ports below 1024 are considered privileged ports, and binding to them typically requires elevated privileges. In many cloud or containerized environments, you should avoid using privileged ports. It's often best to use ports above 1024 to avoid potential permission issues.

3. **Firewall and Security Policies**: If you're deploying your application in a restrictive network environment, make sure the chosen port is allowed through any firewalls or security policies. Port 8080 is a commonly allowed port, but this may not be the case for other ports.

4. **Service Discovery**: If your application relies on service discovery and communicates with other services, ensure that the chosen port aligns with the service discovery configuration.

5. **Environment Variables**: When you choose a port, make sure that your application's configuration, including environment variables, specifies the same port. It's good practice to make the port a configurable setting, allowing it to be easily changed if needed without modifying the code.

Here's an example of how to configure your Go application to use a custom port via an environment variable:

```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not specified in the environment variable.
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, EKS!")
	})

	listenAddress := fmt.Sprintf(":%s", port)
	http.ListenAndServe(listenAddress, nil)
}
```

With this approach, you can set the `APP_PORT` environment variable to specify the desired port when deploying your application. This allows you to easily change the port without modifying the code.


2.)
In the context of configuring network interfaces for your Go application or any networked service, it's important to understand the different options for binding to specific IP addresses and interfaces. The choice of which IP address or interface to use depends on your specific requirements and deployment environment. Here are some considerations:

1. **0.0.0.0 (All Interfaces)**:
   - When you bind to `0.0.0.0`, your application will listen on all available network interfaces. It will accept incoming connections on any network interface, including local loopback (`127.0.0.1`) and any other network interfaces present on the system.
   - This is often used when you want your application to be accessible from any IP address. It's a common choice for web servers that should respond to requests from any available network interface.

2. **Specific IP Address**:
   - You can bind your application to a specific IP address if you want it to listen only on a particular network interface. This can be useful when you have multiple network interfaces on your server, and you want to control which interface your application uses.
   - You should specify the IP address as a string, such as `"127.0.0.1"` for the loopback interface or a specific external IP address.

Here's an example of how to bind your Go application to a specific IP address:

```go
http.ListenAndServe("192.168.1.100:8080", router)
```

In this example, the application will only accept incoming connections on the network interface associated with the IP address `192.168.1.100`.

3. **Localhost (127.0.0.1)**:
   - Binding to `127.0.0.1` means your application will only accept connections from the local machine (loopback). This is used when you want your application to be accessible only from the same server where it's running. It's common for development and debugging purposes.

4. **Dynamic or Configurable IP Addresses**:
   - You can make the choice of IP address dynamic and configurable by reading it from environment variables or configuration files. This allows you to change the binding address without modifying the code. For example:

   ```go
   bindAddress := os.Getenv("APP_BIND_ADDRESS")
   port := os.Getenv("APP_PORT")
   http.ListenAndServe(bindAddress+":"+port, router)
   ```

In production and deployment scenarios, the choice of which IP address and interface to use will depend on factors such as security policies, network configurations, and the desired accessibility of your application. It's important to consider these factors when configuring your application's network interface to ensure it aligns with your specific requirements and network environment.
Obtaining a specific IP address for your project involves several steps, and whether or not you need to purchase an IP address depends on your specific requirements and infrastructure. Here are the common ways to get an IP address:

1. **Public IP Addresses**:

    - **Internet Service Provider (ISP)**: If you're hosting a server or service on your own premises, you can request a public IP address from your Internet Service Provider. This IP address is typically leased to you as part of your internet service plan.

    - **Cloud Providers**: If you're hosting your project on a cloud platform like Amazon Web Services (AWS), Google Cloud Platform (GCP), or Microsoft Azure, you can allocate public IP addresses from the cloud provider's pool of available addresses. Cloud providers often provide a mechanism for requesting and managing public IP addresses through their management consoles or APIs.

2. **Domain Names and DNS**:

    - To make your project accessible via a human-readable domain name (e.g., www.example.com), you need to purchase a domain name from a domain registrar. Domain registrars usually provide domain name registration services and often offer additional services such as DNS hosting.

    - You can associate your domain name with one or more IP addresses through Domain Name System (DNS) records. For example, an "A" record maps a domain name to an IP address. When someone accesses your domain, DNS resolution directs them to the associated IP address.

3. **IPv4 vs. IPv6**:

    - IPv4 addresses are the traditional 32-bit IP addresses, which are running out due to high demand. IPv6 addresses, on the other hand, use a 128-bit format and are designed to provide a vastly larger pool of unique addresses. If you have specific IP requirements and need a large number of IP addresses, consider adopting IPv6.

4. **Static vs. Dynamic IP Addresses**:

    - Public IP addresses can be either static or dynamic. Static IPs do not change over time and are typically used for services that need a fixed, unchanging address. Dynamic IPs may change periodically and are often used for residential internet connections. If you require a static IP, you need to specify this when requesting an IP address.

5. **Port Forwarding and NAT**:

    - In some scenarios, you might not need a public IP address for every server or service. Instead, you can use port forwarding and Network Address Translation (NAT) to map multiple internal IP addresses to a single external IP address. This is a common technique for managing multiple services on a single public IP.

6. **Shared vs. Dedicated IP**:

    - Shared IP addresses are provided to multiple users and can be more cost-effective. Dedicated IP addresses are assigned exclusively to a single user or organization and are often required for specific use cases, such as SSL/TLS certificates.

The specific steps to obtain an IP address and associate it with your project may vary depending on your hosting environment, whether you're using your own servers, cloud services, or a combination of both. In many cases, cloud providers offer public IP address management as part of their service, making it easier to allocate and associate IP addresses with your applications.

Remember that the availability of IP addresses may be subject to regional and provider-specific limitations, and it's important to consider your project's scalability and security requirements when planning for IP address allocation.
*/
