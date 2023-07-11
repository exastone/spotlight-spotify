import "../styles/card.css";

export function AlbumCard() {
    return (
        <div className="card container">
            <div className="album card">
                <div className="album card image">
                    <p>Image</p>
                </div>
                <div className="album card info">
                    <div className="album card info title">
                        <p>Album Title</p>
                    </div>
                    <div className="album card info artist">
                        <p>Artist Name</p>
                    </div>
                </div>
            </div>
        </div>
    )
}