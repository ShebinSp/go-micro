So far we have things working so that it's possible for our user from the front end to hit the broker running in or just send an empty post request
and get some kind of JSON response back then. That's good because we have connectivity, but of course we want to do a lot more than that. So what
we're going to do is add another microservice, we'll add an authentication microservice that exists in Docker beside the broker. And what we want
to do is have the user try to authenticate through the broker. The broker calls the authentication microservice determines whether or not that user
is able to authenticate and sends back the appropriate response.
	        Now, obviously, that means that the authentication microservice has to have some kind of persistence. We need to be able to store user-
information somewhere. So what we'll do is add a Postgres database that's specific to the authentication service. And of course, that means we could
use this authentication service for any kind of application. So if you have multiple services and you have the same users for all of those services,
you could use this microservice for authentication for all of them.