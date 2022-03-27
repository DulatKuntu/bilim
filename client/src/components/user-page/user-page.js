import UserPageTop from "./user-comp/user-page-top";
// import UserPageQuest from "./user-comp/user-page-quest";

import "./user-sass/user-page.sass";

import { useEffect, useState } from "react";

const UserPage = () => {
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
            <div className="user-page">
                <UserPageTop
                    name={card.data.name}
                    surname={card.data.surname}
                    username={card.data.username}
                    interests={card.data.interests}
                />

                <div className="user-page-bio">
                    <div className="user-page-bio__text">{card.data.bio}</div>
                </div>
            </div>
        );
    } else {
        return <h1>Hello world</h1>;
    }
};

export default UserPage;
