import React, { useState, useEffect } from "react";
import "../styles/placeholder.css";

interface Track {
  name: string;
  album: {
    images: [
      { url: string }
    ]
  };
  artists: [
    { name: string }
  ]
}

const track: Track = {
  name: "",
  album: {
    images: [
      { url: "" }
    ]
  },
  artists: [
    { name: "" }
  ]
}

interface WebPlaybackProps {
  token: string;
}

const WebPlayback: React.FC<WebPlaybackProps> = (props) => {

  const [is_paused, setPaused] = useState(false);
  const [is_active, setActive] = useState(false);
  const [player, setPlayer] = useState<any>(undefined);
  const [current_track, setTrack] = useState(track);

  useEffect(() => {

    const script = document.createElement("script");
    script.src = "https://sdk.scdn.co/spotify-player.js";
    script.async = true;

    document.body.appendChild(script);

    window.onSpotifyWebPlaybackSDKReady = () => {

      const player = new (window as any).Spotify.Player({
        name: "Spotlight - Dev App",
        getOAuthToken: (cb: (token: string) => void) => { cb(props.token); },
        volume: 0.5
      });

      setPlayer(player);

      player.addListener("ready", ({ device_id }: { device_id: string }) => {
        console.log("Ready with Device ID", device_id);
      });

      player.addListener("not_ready", ({ device_id }: { device_id: string }) => {
        console.log("Device ID has gone offline", device_id);
      });

      player.addListener("player_state_changed", (state: any) => {

        if (!state) {
          return;
        }

        setTrack(state.track_window.current_track);
        setPaused(state.paused);

        player.getCurrentState().then((state: any) => {
          (!state) ? setActive(false) : setActive(true)
        });

      });

      player.connect();

    };
  }, []);

  if (!is_active) {
    return (
      <div className="container">
        <div className="main-wrapper">
          <b> Instance not active. Transfer your playback using your Spotify app </b>
        </div>
      </div>
    )
  } else {
    return (
      <div className="container">
        <div className="main-wrapper">

          <img src={current_track.album.images[0].url} className="now-playing__cover" alt="" />

          <div className="now-playing__side">
            <div className="now-playing__name">{current_track.name}</div>
            <div className="now-playing__artist">{current_track.artists[0].name}</div>

            <button className="btn-spotify" onClick={() => { player.previousTrack() }} >
              &lt;&lt;
            </button>

            <button className="btn-spotify" onClick={() => { player.togglePlay() }} >
              {is_paused ? "PLAY" : "PAUSE"}
            </button>

            <button className="btn-spotify" onClick={() => { player.nextTrack() }} >
              &gt;&gt;
            </button>
          </div>
        </div>
      </div>
    );
  }
}

export default WebPlayback
