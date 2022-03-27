import React, { useState } from "react";

import "../sign-sass/sign-in.sass";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faEnvelope, faKey } from "@fortawesome/free-solid-svg-icons";

import UserProfilePopUp from "./user-profile-pop-up";

const SignIn = ({ setUserRegistered, setSign, setUserId, user }) => {
    const [submited, setSubmited] = useState(false);
    const [popUp, setPopUp] = useState(false);

    const handleCard = (event) => {
        event.preventDefault();
        setSubmited(true);

        const user = {
            password: event.target.password.value,
            username: event.target.username.value,
        };

        addCardHandler(user);
    };

    let link;
    if (user == "s") {
        link = "http://localhost:4000/unauthed/signin";
    } else if (user == "m") {
        link = "http://localhost:4000/unauthed/signinMentor";
    }

    async function addCardHandler(card) {
        const response = await fetch(
            { link },
            {
                method: "POST",
                body: JSON.stringify(card),
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                    // Accept: "*/*",
                },
            }
        );
        const data = await response.json();
        console.log(data);
        setUserId(data.data.id);
        console.log(typeof data.data.id);
        sessionStorage.setItem("token", data.data.token);
        setUserRegistered(true);
    }

    function getToken() {
        const tokenString = sessionStorage.getItem("token");
        console.log(tokenString);
    }

    return (
        <div className="sign-in">
            <div className="sign-in-welcome">
                Добро пожаловать в
                <br />
                <span className="sign-in-welcome__company">Evision</span>
            </div>

            <div className="sign-in-text">Войти!</div>

            <form className="sign-in-main" onSubmit={handleCard}>
                <FontAwesomeIcon
                    icon={faEnvelope}
                    className="sign-in-main__emoji"
                    id="sign-in-first__emoji"
                />

                <input
                    type="text"
                    name="username"
                    className="sign-in-main__input"
                    placeholder="example@gmail.com"
                />

                <FontAwesomeIcon
                    icon={faKey}
                    className="sign-in-main__emoji"
                    id="sign-in-second__emoji"
                />

                <input
                    type="text"
                    name="password"
                    className="sign-in-main__input"
                    placeholder="Enter your password"
                />

                <input
                    type="submit"
                    name="Password"
                    className="sign-in-main__submit"
                    onClick={() => {
                        setUserRegistered(true);
                    }}
                />
            </form>

            <div className="sign-in-no-account">
                Нету аккаунта?
                <a
                    href="#"
                    className="sign-in-no-account__link"
                    onClick={() => setSign(true)}
                >
                    Зарегестрироваться
                </a>
            </div>

            {submited && <UserProfilePopUp />}
        </div>
    );
};

export default SignIn;
