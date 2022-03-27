import MentorsPageTop from "./mentors-comp/mentors-page-top";
import MentorsPageQuest from "./mentors-comp/mentors-page-quest";

import "./mentors-sass/mentors-page.sass";

import { useEffect } from "react";

const MentorsPage = () => {
    const token = sessionStorage.getItem("token");

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
            console.log(data);
        }

        addCardHandler();
    }, []);

    return (
        <div className="mentors-page">
            <MentorsPageTop />

            <div className="mentors-page-bio">
                <div className="mentors-page-bio__text">
                    Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                    Adipisci nemo nihil neque sed, tempore fugiat ad esse. Ab
                    aliquid aperiam quidem blanditiis cumque assumenda ea sequi
                    doloremque sint. Sapiente, tempore! Lorem ipsum dolor sit
                    amet, consectetur adipisicing elit. Adipisci nemo nihil
                    neque sed, tempore fugiat ad esse. Ab aliquid aperiam quidem
                    blanditiis cumque assumenda ea sequi doloremque sint.
                    Sapiente, tempore!
                </div>
            </div>
            <button className="mentors-page__send">Contact me!</button>
        </div>
    );
};

export default MentorsPage;
