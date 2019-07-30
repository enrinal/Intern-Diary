# Order System Microservices Architecture — A Journey of Understanding

<p align="center">![alt text](https://wwwimg.victorops.com/2015/12/microservices-meme.png)</p>

## What are we build ?
In my [intern-diary repo](https://gitlab.warungpintar.co/enrinal/intern-diary), we already learn how to build order system management, using clean architecture on golang
and in this journey We will be building Order System with Microservice Architecture. I wanted something that could really show-off the separation of complexity. So this felt like a good challenge!

## What is microservice ?
In a traditional monolith application, all of an organisations features are written into one single application. Sometime's they're grouped by their type, such as controllers, entities, factories etc. Other times, perhaps in larger application, features are separated by concern or by feature. So you may have an auth package, a friends package, and an articles package. Which may contain their own set of factories, services, repositories, models etc. But ultimately they are grouped together within a single codebase.

A microservice is the concept of taking that second approach slightly further, and separating those concerns into their own, independent runnable codebase.

## Why microservicess?
Complexity - Splitting features into microservices allows you to split code into smaller chunks. It harks back to the old unix adage of 'doing one thing well'. There's a tendency with monoliths to allow domains to become tightly coupled with one another, and concerns to become blurred. This leads to riskier, more complex updates, potentially more bugs and more difficult integrations.

Scale - In a monolith, certain areas of code may be used more frequently than others. With a monolith, you can only scale the entire codebase. So if your auth service is hit constantly, you need to scale the entire codebase to cope with the load for just your auth service.

With microservices, that separation allows you to scale individual services individually. Meaning more efficient horizontal scaling. Which works very nicely with cloud computing with multiple cores and regions etc.

## Why Golang ?
Microservices are supported by just about all languages, after all, microservices are a concept rather than a specific framework or tool. That being said, some languages are better suited and, or have better support for microservices than others. One language with great support is Golang.

Golang is very light-weight, very fast, and has a fantastic support for concurrency, which is a powerful capability when running across several machines and cores.

Go also contains a very powerful standard library for writing web services.

Finally, there is a fantastic microservice framework available for Go called go-micro. Which we will be using in this series.
