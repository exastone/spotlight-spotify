import { useEffect } from "react";

interface WebPlayerV2Props {
  token: string;
}

const WebPlayerV2: React.FC<WebPlayerV2Props> = (token) => {

  useEffect(() => {
    // Load the Spotify Web Playback SDK script dynamically
    const scriptTag = document.createElement('script');
    scriptTag.src = 'https://sdk.scdn.co/spotify-player.js';
    scriptTag.async = true;
    document.body.appendChild(scriptTag);

    // The callback function that will be called when the SDK is ready
    window.onSpotifyWebPlaybackSDKReady = () => {
      const token = '[My access token]'; // Replace with your access token

      // Initialize the player
      const player = new Spotify.Player({
        name: 'Web Playback SDK Quick Start Player',
        getOAuthToken: (cb) => {
          cb(token);
        },
      });

      // Add any event listeners and other player logic here
    };
  }, []);


  return (
    <div>
      <h1>Spotify Web Playback SDK Quick Start</h1>
      {/* Add any other JSX elements for your player UI */}
    </div>
  );
}
