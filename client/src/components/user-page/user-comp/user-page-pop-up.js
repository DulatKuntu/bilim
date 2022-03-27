import React, { useState } from "react";

import "../user-sass/user-page-pop-up.sass";

import { Link } from "react-router-dom";

const UserPagePopUp = ({ setPopUp }) => {
    const handleCard = (event) => {
        event.preventDefault();
        // setSubmited(true);

        const user = {};

        if (event.target.checkbox_1.checked) {
            user["1"] = 1;
        }
        if (event.target.checkbox_2.checked) {
            user["2"] = 2;
        }
        if (event.target.checkbox_3.checked) {
            user["3"] = 3;
        }
        if (event.target.checkbox_4.checked) {
            user["4"] = 4;
        }

        console.log(user);

        addCardHandler(user);
    };

    async function addCardHandler(card) {
        const response = await fetch(
            "http://localhost:4000/authed/addInterest",
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
    }

    return (
        <div className="user-page-pop-up">
            <div className="user-page-pop-up__container">
                <h2 className="user-page-pop-up__text">Выберите интерес:</h2>

                <form
                    className="user-page-pop-up__container_main"
                    onSubmit={handleCard}
                >
                    <div className="user-page-pop-up__container_main_element">
                        <input
                            type="checkbox"
                            name="checkbox_1"
                            id=""
                            className="user-page-pop-up__container_main_element_input"
                            value="1"
                        />
                        <div className="user-page-pop-up__container_main_element_text">
                            Программирование
                        </div>
                    </div>
                    <div className="user-page-pop-up__container_main_element">
                        <input
                            type="checkbox"
                            name="checkbox_2"
                            id=""
                            className="user-page-pop-up__container_main_element_input"
                            value="2"
                        />
                        <div className="user-page-pop-up__container_main_element_text">
                            Математика
                        </div>
                    </div>
                    <div className="user-page-pop-up__container_main_element">
                        <input
                            type="checkbox"
                            name="checkbox_3"
                            id=""
                            className="user-page-pop-up__container_main_element_input"
                            value="3"
                        />
                        <div className="user-page-pop-up__container_main_element_text">
                            Биология
                        </div>
                    </div>
                    <div className="user-page-pop-up__container_main_element">
                        <input
                            type="checkbox"
                            name="checkbox_4"
                            id=""
                            className="user-page-pop-up__container_main_element_input"
                            value="4"
                        />
                        <div className="user-page-pop-up__container_main_element_text">
                            Физика
                        </div>
                    </div>

                    <button className="user-page-pop-up__container_main__submit">
                        Выбрать
                    </button>
                </form>

                <button
                    className="user-page-pop-up__container__close"
                    onClick={() => {
                        setPopUp(false);
                    }}
                >
                    X
                </button>
            </div>
        </div>
    );
};

export default UserPagePopUp;
