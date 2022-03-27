import "./mentors-sass/mentors.sass";

import MentorsCard from "./mentors-comp/mentors-card";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch } from "@fortawesome/free-solid-svg-icons";

import { Link } from "react-router-dom";
import { useState, useEffect } from "react";

const Mentors = () => {
    const token = sessionStorage.getItem("token");

    const [card, SetCard] = useState(null);

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
            console.log(data);
            SetCard(data);
        }

        addCardHandler();
    }, []);

    if (card) {
        console.log(card);
        return (
            <div className="mentors">
                <div className="mentors-top">
                    <div className="mentors-top-text">mentors</div>

                    <form className="mentors-top-search">
                        <FontAwesomeIcon
                            icon={faSearch}
                            className="mentors-top-search__emoji"
                        />
                        <input
                            type="text"
                            className="mentors-top-search__input"
                            placeholder="Search among 1000+ mentors"
                        />
                    </form>
                </div>

                <div className="mentors-container">
                    <Link to="/mentors/1" style={{ textDecoration: "none" }}>
                        <MentorsCard />
                    </Link>
                    <MentorsCard />
                    <MentorsCard />
                </div>
            </div>
        );
    } else {
        return <h1>Hello world</h1>;
    }
};

export default Mentors;
