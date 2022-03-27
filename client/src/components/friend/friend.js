import React, { useEffect, useState } from "react";

import "./friend.sass";
import MentorsCard from "../mentors/mentors-comp/mentors-card";

const Friend = () => {
    const token = sessionStorage.getItem("token");

    const [card, setCard] = useState([]);

    useEffect(() => {
        async function addCardHandler(card) {
            const response = await fetch(
                "http://localhost:4000/authed/getMentors",
                {
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded",
                        // Accept: "*/*",
                        Authorization: `${token}`,
                    },
                }
            );
            const data = await response.json();
            setCard(data.data);
        }

        addCardHandler();
    }, []);

    if (card) {
        return (
            <div className="friend">
                <div className="friend-top">
                    <h2 className="friend-top-header">Найди себе друга!</h2>
                    <button className="friend-top-create">
                        Создать запрос
                    </button>
                </div>
                <div className="friend-container">
                    {card.map((el) => {
                        return (
                            <MentorsCard
                                name={el.name}
                                interests={el.interests}
                                bio={el.bio}
                            />
                        );
                    })}
                </div>
            </div>
        );
    } else {
        return <h1>Hello world</h1>;
    }
};

export default Friend;
