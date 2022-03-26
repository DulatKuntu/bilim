import React from "react";

import "../sign-sass/sign-up.sass";

const SignIn = ({ setSign }) => {
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

    async function addCardHandler(card) {
        const response = await fetch("http://localhost:8080/api/task/create", {
            method: "POST",
            body: JSON.stringify(card),
            headers: {
                // "Content-Type": "application/card",
                Accept: "*/*",
            },
        });
        const data = await response.json();
        console.log(data);
    }

    return (
        <div className="sign-in">
            <div className="sign-in-welcome">
                Welcome to
                <br />
                <span className="sign-in-welcome__company">Company name</span>
            </div>

            <div className="sign-in-text">Sign up!</div>

            <form className="sign-in-main" onSubmit={handleCard}>
                <div className="sign-in-main__inputs">
                    <div className="sign-in-main__inputs_block">
                        <div className="sign-in-main__inputs_block_element">
                            <label htmlFor="name">Enter your name!</label>
                            <input
                                type="text"
                                name="name"
                                className="sign-in-main__inputs_block_element__input"
                                placeholder="here!"
                            />
                        </div>

                        <div className="sign-in-main__inputs_block_element">
                            <label htmlFor="surname">Enter your surname!</label>
                            <input
                                type="text"
                                name="surname"
                                className="sign-in-main__inputs_block_element__input"
                                placeholder="here!"
                            />
                        </div>
                    </div>

                    <div className="sign-in-main__inputs_block">
                        <div className="sign-in-main__inputs_block_element">
                            <label
                                htmlFor="username"
                                className="sign-in-main__inputs_block_element_label"
                            >
                                Enter your username!
                            </label>
                            <input
                                type="text"
                                name="username"
                                className="sign-in-main__inputs_block_element__input"
                                placeholder="here!"
                            />
                        </div>

                        <div className="sign-in-main__inputs_block_element">
                            <label
                                htmlFor="contacts"
                                className="sign-in-main__inputs_block_element_label"
                            >
                                Enter your contacts!
                            </label>
                            <input
                                type="text"
                                name="contacts"
                                className="sign-in-main__inputs_block_element__input"
                                placeholder="here!"
                            />
                        </div>
                    </div>

                    <div className="sign-in-main__inputs_block">
                        <div className="sign-in-main__inputs_block_element">
                            <label
                                htmlFor="email"
                                className="sign-in-main__inputs_block_element_label"
                            >
                                Enter your email!
                            </label>
                            <input
                                type="text"
                                name="email"
                                className="sign-in-main__inputs_block_element__input"
                                placeholder="example@gmail.com"
                            />
                        </div>

                        <div className="sign-in-main__inputs_block_element">
                            <label
                                htmlFor="password"
                                className="sign-in-main__inputs_block_element_label"
                            >
                                Enter your password!
                            </label>
                            <input
                                type="text"
                                name="password"
                                className="sign-in-main__inputs_block_element__input"
                                placeholder="**********"
                            />
                        </div>
                    </div>
                </div>

                <label htmlFor="bio" className="sign-in-main__bio_label">
                    Enter your bio!
                </label>
                <input
                    type="text"
                    name="bio"
                    className="sign-in-main__bio"
                    placeholder="here!"
                />

                <input
                    type="submit"
                    name="Password"
                    className="sign-in-main__submit"
                />
            </form>

            <div className="sign-in-no-account">
                Don't have an account?
                <a
                    href="#"
                    className="sign-in-no-account__link"
                    onClick={() => setSign(true)}
                >
                    Register
                </a>
            </div>
        </div>
    );
};

export default SignIn;
