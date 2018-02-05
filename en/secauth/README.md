# Security, Authentication, Authorization

# Securing your applications

Whether you deploy in the cloud or on-premises, the security of your application is critical to your business success.  Nothing sinks a ship faster than a breach.

In this module we'll talk about the various ways to make your application more secure.

# Ensuring valid communications between systems

If attackers gain entry to your network, they have access to any system that can be reached from their entry point.

One way to mitigate this attack vector is to validate the identity of services that communicate with each other. 

There are a several ways to validate services.

- Client certificates
- Vault from Hashicorp

# Client Certificates

Client certificates are a great way to ensure that the communications between services are valid.  They require a little bit of infrastructure, however, and planning.

- Must create own Certificate Authority, and deploy the CA trust keys to every server
- This isn't something you want to do by hand

[CloudFlare cfssl](https://github.com/cloudflare/cfssl) 

# Vault from Hashicorp

[Vault](https://www.vaultproject.io) 

Vault secures, stores, and tightly controls access to tokens, passwords, certificates, API keys, and other secrets in modern computing. Vault handles leasing, key revocation, key rolling, and auditing. Through a unified API, users can access an encrypted Key/Value store and network encryption-as-a-service, or generate AWS IAM/STS credentials, SQL/NoSQL databases, X.509 certificates, SSH credentials, and more.

# Recommendation

Use Vault, and start from Day 1.  

	It's harder to retrofit security than it is to plan for it from the beginning.

# Transport security

No network is secure.  Ever.  That should be your base assumption when going into any security planning.

There are things you can do to decrease your risk.

- TLS between all services
- Kubernetes/Docker -- Use Romana, OpenContrail, Calico to provide network policies on your container substrate

[Romana Slides](http://www.slideshare.net/RomanaProject/kubecon-london-2016-ronana-cloud-native-sdn) 
 
# Authentication

Authentication is the act of verifying the identity of an entity interacting with your system.

- Username/Password
- Unique token

Doesn't only apply to people. Machines and services can and should be authenticated too.

	Use Vault For Machines And Services

### Users:

If possible, don't roll your own. Use LDAP, or Active Directory, or a central authorization system like CoreOS Dex

[DEX](https://github.com/coreos/dex)

# Authorization

Authorization is the act of verifying the permissions of an entity as it relates to an action they wish to perform.

- Often a Role based system.  Users belong to roles, roles have permissions to perform an activity

[Go-Micro Authz/n](https://github.com/micro/auth-srv)

# Tokens and claims

Tokens and claims help you reduce cost of security.

JWT Tokens provide both authentication and authorization with "Claims"

[JWT specification](https://jwt.io)

Authorization claims are embedded into the token, and signed, providing a cryptographically secure way for the user/service to send both their authentication and authorization with every request.

	{
		"sub": "1234567890", // public claim
		"name": "John Doe", // public claim
		"role": "admin"// PRIVATE CLAIM - YOU create this part
	}

JWT Tokens are represented in JSON format before encryption, so you can make complex claims.  i.e. Embed a complex json document as a private claim.

# JWT Tools in Go

Canonical implementation:  
- [JWT-go](https://github.com/dgrijalva/jwt-go)

# Exercise 1

Run the Vault getting started guide on a single node.  Skip the AWS credentials part.

[Vault](https://www.vaultproject.io) 

# Exercise 2
A JWT Implementation from a previous project:

[congo](https://github.com/gopheracademy/congo) 

Walk through the code together to explore how JWT works.
