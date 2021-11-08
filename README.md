# 📖 Tutorial: Asynq. Simple, reliable & efficient distributed task queue for your next Go project

Asynq is a Go library for queueing tasks and processing them asynchronously with workers. It's backed by Redis and is designed to be scalable yet easy to get started. Task queues are used as a mechanism to distribute work across multiple machines.

👉 The full article is published on **April 7, 2021**, on Dev.to: https://dev.to/koddr/asynq-simple-reliable-efficient-distributed-task-queue-for-your-next-go-project-4jhg

![article preview pic](https://user-images.githubusercontent.com/11155743/113790679-89c9a600-974a-11eb-9aca-b4d15187e6cb.jpg)

## Quick start

1. Clone this repository and go to it root folder.
2. Start a Redis server (by Docker or locally).
3. Start Asynq worker server:

```console
make worker
```

4. Start generating tasks by Asynq client:

```console
make client
```

5. Install [Asynqmon](https://github.com/hibiken/asynqmon) (Asynq web UI) to your system.
6. Go to [localhost:8080](http://localhost:8080) and see:

![Screenshot](https://user-images.githubusercontent.com/11155743/113557216-57af2b80-9606-11eb-8ab6-df023b14e5c1.png)

## P.S.

If you want more articles like this on this blog, then post a comment below and subscribe to me. Thanks! 😘

And, of course, you can support me by donating at [LiberaPay](https://liberapay.com/koddr/donate). _Each donation will be used to write new articles and develop non-profit open-source projects for the community._

[![Support author at LiberaPay](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/zq8442cqyjq2i1jdeay8.png)](https://liberapay.com/koddr/donate)

## ⚠️ License

MIT &copy; [Vic Shóstak](https://shostak.dev/) & [True web artisans](https://1wa.co/).
