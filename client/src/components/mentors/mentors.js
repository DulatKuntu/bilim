import "./mentors-sass/mentors.sass";

import MentorsCard from "./mentors-comp/mentors-card";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch } from "@fortawesome/free-solid-svg-icons";

import { Link } from "react-router-dom";
import { useState, useEffect } from "react";

const Mentors = ({ card, setCard }) => {
    const token = sessionStorage.getItem("token");

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
        console.log("Hello");
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
                    {card.map((el) => {
                        return (
                            <Link
                                to={`/mentors/${el.id}`}
                                style={{ textDecoration: "none" }}
                            >
                                <MentorsCard
                                    name={el.name}
                                    interests={el.interests}
                                    bio={el.bio}
                                />
                            </Link>
                        );
                    })}
                </div>
            </div>
        );
    } else {
        return <h1>Hello world</h1>;
    }
};

export default Mentors;
