import React from 'react'

type Props = {
    image: string;
    title: string;
    artist: string;
    type: string;
};

const ListItem: React.FC<Props> = ({ image, title, artist, type }) => {
    return (
        <li className="library list-item">
            <img src={image} alt={title} />
            <div className="col">
                <span className="title">
                    {title}
                </span>
                <span className="artist">
                    {type + " "} • {" " + artist}
                </span>
            </div>
        </li>
    )
}

export default ListItem