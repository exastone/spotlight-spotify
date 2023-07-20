import { useState } from "react";

import { CurrentlyPlaying } from "./CurrentlyPlaying";
import { OtherControls } from "./OtherControls";
import { PlaybackControls } from "./PlaybackControls";

import "../styles/player.css"

const track = {
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

export function Player() {
    const [isPaused, setIsPaused] = useState(false);

    /* Is active is used to determine if the current playback is set to the app;
     i.e. not active, transfer playback to the app using a spotify app */
    const [isActive, setIsActive] = useState(false);

    const [currentTrack, setCurrentTrack] = useState(track);

    return (
        <div className="player">
            <div className="player-item item-1">
                <CurrentlyPlaying />
            </div>
            <div className="player-item item-2">
                <PlaybackControls />
            </div>
            <div className="player-item item-3">
                <OtherControls />
            </div>
        </div>
    );
}