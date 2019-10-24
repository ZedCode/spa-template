# spa-template
This is an example project that aims to give a starting point for creating a new single page web application. It has a simple to expand on folder structure with a few useful reminders about how to bind/load data with Knockout.js.

## Working with this example
To start working with this example, clone it, rename the folder, delete the `.git` directory and start making changes. The goal of a SPA is that you only load the index page once, all subsequent data is loaded via AJAX calls to the API. The jquery AJAX calls are then bound to Knockout.js observables to be dynamically updated.

Below is a description of what you will find in each location of this repository and what it's for.

### static
This directory contains a `js` and a `css` folder. Each of these contains the relevant Bootstrap, Knockout, jQuery and Popper.js. There is a route in `vendor/spaserver/server.go` called `staticServer` which is responsible for serving this content. The static server route can be expanded to handle images and other arbitrary file types.

### templates
The `templates` directory is used by the (Render Library)[https://github.com/unrolled/render] to search for template files by name. The primary template is `index.tmpl` and includes the other templates here.

### vendor
The vendor directory contains the `spaserver` library. This is where you would expand the kind of data you return, the API routes that are available, the methods you can use for each route and so on. There is an example included for both inserting data into the template at execution time (for example, if you want your SPA to be configurable, but not change once the application is up, you would load those things via the template engine). There is also an example for retreiving data over the API route when you press the `get data` button.

## Building
You can use the included Makefile to produce a Docker image you can then run and browse to on your local computer. Build and run as:
```
$ make linux && make run
```
You can then browse to localhost:8080 to interact with the demo.