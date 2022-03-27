import React, { useState } from "react";

import "../sign-sass/sign-up.sass";

import { Link } from "react-router-dom";

import UserProfilePopUp1 from "./user-profile-pop-up1";

const SignUp = ({ setSign, user }) => {
    const [submited, setSubmited] = useState(false);

    const handleCard = (event) => {
        event.preventDefault();

        const user = {
            name: event.target.name.value,
            surname: event.target.surname.value,
            username: event.target.username.value,
            contacts: event.target.contacts.value,
            password: event.target.password.value,
            email: event.target.email.value,
            bio: event.target.bio.value,
        };

        addCardHandler(user);
    };

    let link;
    if (user == "s") {
        link = "http://localhost:4000/unauthed/signup";
    } else if (user == "m") {
        link = "http://localhost:4000/unauthed/signupMentor";
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
    }

    return (
        <div className="sign-up">
            <div className="sign-up-welcome">
                Welcome to
                <br />
                <span className="sign-up-welcome__company">Company name</span>
            </div>

            <div className="sign-up-text">Sign up!</div>

            <form className="sign-up-main" onSubmit={handleCard}>
                <div className="sign-up-main__inputs">
                    <div className="sign-up-main__inputs_block">
                        <div className="sign-up-main__inputs_block_element">
                            <label htmlFor="name">Enter your name!</label>
                            <input
                                type="text"
                                name="name"
                                className="sign-up-main__inputs_block_element__input"
                                placeholder="here!"
                            />
                        </div>

                        <div className="sign-up-main__inputs_block_element">
                            <label htmlFor="surname">Enter your surname!</label>
                            <input
                                type="text"
                                name="surname"
                                className="sign-up-main__inputs_block_element__input"
                                placeholder="here!"
                            />
                        </div>
                    </div>

                    <div className="sign-up-main__inputs_block">
                        <div className="sign-up-main__inputs_block_element">
                            <label
                                htmlFor="username"
                                className="sign-up-main__inputs_block_element_label"
                            >
                                Enter your username!
                            </label>
                            <input
                                type="text"
                                name="username"
                                className="sign-up-main__inputs_block_element__input"
                                placeholder="here!"
                            />
                        </div>

                        <div className="sign-up-main__inputs_block_element">
                            <label
                                htmlFor="contacts"
                                className="sign-up-main__inputs_block_element_label"
                            >
                                Enter your contacts!
                            </label>
                            <input
                                type="text"
                                name="contacts"
                                className="sign-up-main__inputs_block_element__input"
                                placeholder="here!"
                            />
                        </div>
                    </div>

                    <div className="sign-up-main__inputs_block">
                        <div className="sign-up-main__inputs_block_element">
                            <label
                                htmlFor="email"
                                className="sign-up-main__inputs_block_element_label"
                            >
                                Enter your email!
                            </label>
                            <input
                                type="text"
                                name="email"
                                className="sign-up-main__inputs_block_element__input"
                                placeholder="example@gmail.com"
                            />
                        </div>

                        <div className="sign-up-main__inputs_block_element">
                            <label
                                htmlFor="password"
                                className="sign-up-main__inputs_block_element_label"
                            >
                                Enter your password!
                            </label>
                            <input
                                type="text"
                                name="password"
                                className="sign-up-main__inputs_block_element__input"
                                placeholder="**********"
                            />
                        </div>
                    </div>
                </div>

                <label htmlFor="bio" className="sign-up-main__bio_label">
                    Enter your bio!
                </label>
                <input
                    type="text"
                    name="bio"
                    className="sign-up-main__bio"
                    placeholder="here!"
                />

                {/* <Link to="/sign/survey"> */}
                <input
                    type="submit"
                    name="Password"
                    className="sign-up-main__submit"
                    onClick={() => {
                        setSubmited(true);
                    }}
                />
                {/* </Link> */}
            </form>

            <div className="sign-up-no-account">
                Have an account?
                <a
                    href="#"
                    className="sign-up-no-account__link"
                    onClick={() => setSign(false)}
                >
                    Go back
                </a>
            </div>

            {submited && <UserProfilePopUp1 user={user} />}
        </div>
    );
};

export default SignUp;
