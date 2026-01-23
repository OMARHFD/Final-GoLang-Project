## in this project the main goals will be : 

1- Understanding how to use the docs and look for an information that will last longer in my mind <br >
2- Tackling a chalenge without having a clear vision already <br >
3- Implementing the GoLang properties and mastering it <br >
4- Understanding other subjects included in this project <br >


## in this project the rules are :

1- Using 0/1e9 of AI literally no ai even for comments or readme or git or code obviously <br >
2- Not being afraid after reading something that makes non sense for me  <br >

## Explanation :
### Part 1 : Reverse Proxy
#### Goal :
to make things crystal clear , let's supppose that we have a web application hosted in many backend servers , if some user wants to browse or use our app , we should think about giving him the best possible experience with the ressources we have , the ressources here are our 'many' backemd servers , that s trhe job we want to do using our reverse proxy.
#### NB :  A REVERSE PROXY MAY HAVE MANY OTHER FUNCTIONNALITIES (OTHERS THAN A LOAD BALANCER) AS YOU CAN SEE WE CAN CHANGE THE RESPONSE OF ONE OF OUR BACKEND SERVERS , BBUT OUR MAIN GOAL IN THIS PROJECT IS THE LOAD BALANCING FUNCTIONNALITY.
####Implementation :
The main.go part:
at first we initialized backends in our main function , we could have used a json file , and we will when we will implement the ADMIN API , then after that we created a serverpool because throught it we will use the Round Robin algorithm and then we created our hero -the reverse proxy - it i sa struct in the gho language in the package httputil , that struct has a Dirtector function that takes the request and operater on it (in place) then executing it or forwarding it , when forwarding we used the backend server that we got using our Round Robin algorithm. the reverse proxy struct contains also ModifyResponse function it is an optionnal function that operate on the result of the backend server (it is locate d between the response of the proxy and the response received by the client) in our case we are just printing to the terminal ,so nothing changed. and then finally we launched a server and used our reverseproxy as our handler to execute all these things.
The backend folder part:
you may realise that in our Director function we either forward the request to (we set it s url) one of local hosts 8081...8083 , obviously they should be launched and tha s all of our code in any of backend{ 1 or 2 or 3}.go , in adition we write a message to the respoinseWriter to know what server responded , and obviously to test our Round Robin algorithm , and for fun we prionted the nmessage to the terminal.
