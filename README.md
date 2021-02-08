# helmAPI
A simple API to automate deployment and management of Helm charts within a cluster

This API is inspired from [helm-web-api](https://github.com/microsoft/helm-web-api), which doesn't support Helm v3. This API is backwards compatible with helm-web-api, so requests that worked with helm-web-api should also work with this one.


## Documentation
You can read the documentation [here](https://documenter.getpostman.com/view/7024275/TW76C4SM#5680197b-199a-4f3b-8f8e-2f02fc30ab8a)

## Why does this even exist?
Good question. If you have strong automation needs for kubernetes, by all means use k8s operators (check out the [operator framework](https://operatorframework.io/)) or something. But if you're like me and your automation needs are simple (or maybe you already have a lot of stuff written as helm charts), this API is a quick solution. 
