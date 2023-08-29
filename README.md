# Spotlite - Spotify Client

## Why a Spotify Client?

- I want light mode. Light mode has been a requested feature of Spotify for [nearly a decade](https://community.spotify.com/t5/Live-Ideas/All-Platforms-Light-Mode-option/idi-p/730341), clearly they don't have intentions of bringing light mode to the app so I'll do it myself.

- Spotify's web and desktop client's vary significantly from the mobile (iOS) app and are worst-off because of it, such as a lack of a *What's New* feed on the on mobile, inability to hide audiobooks, poor control of listing history, etc.

### Current State

- UI is framed up and work to connect the frontend to the Spotify Web API is underway
- OAuth was a tedious effort to get working on the backend but authorization code flow now works! (you'll be redirected to a temporary frontend until I get the new frontend connected) With the addition of the new [TypeScript SDK for the Spotify Web API](https://developer.spotify.com/blog/2023-07-03-typescript-sdk) there might be a overhauled
- Current focus is on getting the UI components connected to provide a usable (albeit, minimal functionality) good looking client, before tacking larger tasks such as asset and data caching.

#### UI Preview

![Spotlite - Light mode](spotlight-light.png)
![Spotlite - Dark mode](spotlight-dark.png)
