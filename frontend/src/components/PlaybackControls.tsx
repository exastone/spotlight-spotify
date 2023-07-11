import { useState } from "react";
import { CtrlPrevSong } from "./CtrlPrevSong";
import { CtrlShuffle } from "./CtrlShuffle";
import { CtrlNextSong } from "./CtrlNextSong";
import { CtrlRepeat } from "./CtrlRepeat";
import CtrlPlayPause from "./CtrlPlayPause";
import { Line } from "rc-progress";

export function PlaybackControls() {
    const [isPlaying, setIsPlaying] = useState(false);
    const togglePlayPause = () => {
        setIsPlaying(!isPlaying);
    };

    return (
        <div className="playback-controls">
            <div className="controls-group">
                <button className="control shuffle">
                    <CtrlShuffle />
                </button>
                <button className="control previous-song">
                    <CtrlPrevSong />
                </button>
                <button className="control play-pause" onClick={togglePlayPause}>
                    <CtrlPlayPause isPlaying={isPlaying} />
                </button>
                <button className="control next-song">
                    <CtrlNextSong />
                </button>
                <button className="control repeat">
                    <CtrlRepeat />
                </button>
            </div>
            <div className="playback-progress">
                <div className="playback-item current-time">Curr. time</div>
                <Line className="playback-item progress-bar" percent={30} strokeWidth={1.25} trailWidth={1} />
                <div className="playback-item total-time">Totl. time</div>
            </div>
        </div>
    );
}