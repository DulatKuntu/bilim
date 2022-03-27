import "../user-sass/user-page-top.sass";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faEnvelope, faPhoneAlt } from "@fortawesome/free-solid-svg-icons";

import { useState } from "react";

import UserPagePopUp from "./user-page-pop-up";

import img1 from "../../img/user/1.png";
import img2 from "../../img/user/2.png";
import img3 from "../../img/user/3.png";
import img4 from "../../img/user/4.png";
import img5 from "../../img/user/5.png";
import img6 from "../../img/user/6.png";
import img7 from "../../img/user/7.png";
import img8 from "../../img/user/8.png";
import img9 from "../../img/user/9.png";
import img10 from "../../img/user/10.png";

const UserPageTop = ({ name, surname, username, interests }) => {
    const [popUp, setPopUp] = useState(false);

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

    return (
        <div className="user-page-top">
            <div className="user-page-top-information">
                <div className="user-page-top-information__img_container">
                    <img
                        src={randomImage}
                        className="user-page-top-information__img"
                    />
                </div>
            </div>

            <div className="user-page-top-information__main">
                <div className="user-page-top-information__main_name">
                    {`${name} ${surname}`}
                </div>
                <div className="user-page-top-information__main_username">
                    {username}
                </div>
                <div className="user-page-top-information__main_interests">
                    {interests}
                </div>
            </div>

            <button
                className="user-page-top__select_interest"
                onClick={() => {
                    setPopUp(true);
                }}
            >
                Выберите ваши интересы
            </button>

            {popUp && <UserPagePopUp setPopUp={setPopUp} />}
        </div>
    );
};

export default UserPageTop;
