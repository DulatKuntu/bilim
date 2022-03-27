import MentorsPageTop from "./mentors-comp/mentors-page-top";
import MentorsPageQuest from "./mentors-comp/mentors-page-quest";

import "./mentors-sass/mentors-page.sass";

import { useEffect, useState } from "react";

const MentorsPage = () => {
    const token = sessionStorage.getItem("token");

    const [card, SetCard] = useState(null);

    useEffect(() => {
        async function addCardHandler(card) {
            const response = await fetch(
                "http://localhost:4000/authed/getProfile",
                {
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded",
                        // Accept: "*/*",
                        Authorization: `${token}`,
                    },
                }
            );
            const data = await response.json();
            SetCard(data);
        }

        addCardHandler();
    }, []);

    if (card) {
        console.log(card);

        return (
            <div className="mentors-page">
                <MentorsPageTop
                    name={card.data.name}
                    surname={card.data.surname}
                    username={card.data.username}
                    interests={card.data.interests}
                />

                <div className="mentors-page-bio">
                    <div className="mentors-page-bio__text">
                        {card.data.bio}
                    </div>
                </div>
                <button className="mentors-page__send">Contact me!</button>
            </div>
        );
    } else {
        return <h1>Hello world</h1>;
    }
};

export default MentorsPage;
