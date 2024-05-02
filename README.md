# Hostless-Echo-Example


The Echo project is a powerful and versatile web framework for building scalable and high-performance web applications in the Go programming language. It follows the principles of simplicity, flexibility, and performance to provide developers with an efficient toolkit for building robust web applications. 

Check out their official documentation [here](https://echo.labstack.com/docs/quick-start)

## Deploy a Echo App

### Deployment Instructions

In this tutorial, we advice creating a Dockerfile and using Docker as the build system

1. Fork/Clone this [Hostless-Echo-Example](https://github.com/Hostless-Examples/Hostless-Echo-Example.git) repo from github
2. Click on 'Create New App'
3. Choose a suitable app name
4. Choose your github account
5. Choose the forked github repo/the cloned remote repo
6. Choose a build system

    1. 'Detect automatically with Nixpicks' - automatically detects the programming language and builds the app
    2. 'Docker' - looks for a Dockerfile in the root of the project and build based on the instructions

7. Input a start command(optional)
8. The PORT environment variable is set by Hostless

#### Sample configuration
![sample](https://res.cloudinary.com/do58rrxug/image/upload/v1714681675/Screenshot_2024-05-02_at_21.27.31_ubnwln.png)

#### Example project
An example project is hosted on [https://hostless-echo-example.hostless.app/](https://hostless-echo-example.hostless.app/)

