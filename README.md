# ukiyo

## Idea incubation

- ukiyo will act as a watcher for docker containers. It will run alongside with the other running containers and will be responsible for automatic updates. Updates will be based on push based model compared to existing solutions such as [watchtower](https://github.com/containrrr/watchtower) and [ouroboros](https://github.com/pyouroboros/ouroboros)

- Push events will be recived from ukiyo via webhooks. Docker registries provide webhooks to subscribe and listen to image changes. Locally running images will change only after such an event is received by ukiyo.

- Pull based model can be implemented as an optional way of updating the running containers.

## Components

- Container manager
- Push manager (webhooks configuration) 
- OPTIONAL - Pull based update implementation
