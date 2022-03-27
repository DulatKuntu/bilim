import MentorsPageTop from "./mentors-comp/mentors-page-top";
import MentorsPageQuest from "./mentors-comp/mentors-page-quest";

import "./mentors-sass/mentors-page.sass";

import { useEffect, useState } from "react";

import { useParams } from "react-router-dom";

import img1 from "../img/user/2.png";
import img2 from "../img/user/2.png";
import img3 from "../img/user/3.png";
import img4 from "../img/user/4.png";
import img5 from "../img/user/5.png";
import img6 from "../img/user/6.png";
import img7 from "../img/user/7.png";
import img8 from "../img/user/8.png";
import img9 from "../img/user/9.png";
import img10 from "../img/user/10.png";

const MentorsPage = ({ card }) => {
    const images = [
        img1,
        img2,
        img3,
        img4,
        img5,
        img6,
        img7,
        img8,
        img9,
        img10,
    ];

    const randomImage = images[Math.floor(Math.random() * 10)];

    const token = sessionStorage.getItem("token");

    let { id } = useParams();

    console.log(card);

    const elem = card.filter((el) => {
        return el.id == id;
    });

    if (elem) {
        console.log(elem);

        return (
            <div className="mentors-page">
                <MentorsPageTop
                    name={elem[0].name}
                    surname={elem[0].surname}
                    username={elem[0].username}
                    interests={elem[0].interests}
                    img={randomImage}
                />

                <div className="mentors-page-bio">
                    <div className="mentors-page-bio__text">{elem[0].bio}</div>
                </div>
                <button className="mentors-page__send">Contact me!</button>
            </div>
        );
    } else {
        return <h1>Hello world</h1>;
    }
};

export default MentorsPage;
