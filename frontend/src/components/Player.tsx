import "../styles/player.css"
import { CurrentlyPlaying } from "./CurrentlyPlaying";
import { OtherControls } from "./OtherControls";
import { PlaybackControls } from "./PlaybackControls";

export function Player() {
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