# Twitter Cleaner

Automatically delete tweets, retweets, and favorites from your timeline.

## Usage

### Authentication

This project uses the OAuth 1.0 API, and as such requires api keys to authenticate itself against the Twitter API.

In order to generate those API keys, you'll need to create a new [twitter app][twitter-new-app].

> Note: Please, make sure you change the authentication settings and enable the OAuth 1.0 API with "Read and Write" access.

### Run

**Please, make sure you really want to delete all your tweets before running the available commands below**.

You can run this project using the available docker image as below.

Providing **CLI flags**:
```bash
docker run ghcr.io/hpedrorodrigues/twitter-cleaner \
  -consumer-key <redacted> \
  -consumer-secret <redacted> \
  -access-token <redacted> \
  -access-token-secret <redacted>
```

Providing **environment variable**:
```bash
docker run \
  -e TWITTER_CONSUMER_KEY=<redacted> \
  -e TWITTER_CONSUMER_SECRET=<redacted> \
  -e TWITTER_ACCESS_TOKEN=<redacted> \
  -e TWITTER_ACCESS_TOKEN_SECRET=<redacted> \
  ghcr.io/hpedrorodrigues/twitter-cleaner
```


[twitter-new-app]: https://developer.twitter.com/apps/new
