# 📖 Tutorial: Asynq. Simple, reliable & efficient distributed task queue for your next Go project

👉 The full article is published on **April 7, 2021**, on Dev.to:

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

## ⚠️ License

MIT &copy; [Vic Shóstak](https://shostak.dev/) & [True web artisans](https://1wa.co/).
