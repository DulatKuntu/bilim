import React from "react";

import "../header-sass/header-nav.sass";

import { Link } from "react-router-dom";

import logo from "../header-img/logo.png";

const HeaderNav = ({ userRegistered, userId, user }) => {
    return (
        <div className="header-nav">
            <div className="header-nav-logo__container">
                <img src={logo} alt="" className="header-nav-logo" />
            </div>

            <div className="header-nav-company">EVISION</div>

            {user == "s" && (
                <ul className="header-nav-main">
                    <li className="header-nav-main__elements">
                        <Link
                            to="/mentors"
                            className="header-nav-main__elements_link"
                        >
                            Найти ментора
                        </Link>
                    </li>
                    <li className="header-nav-main__elements">
                        <Link
                            to="/friend"
                            className="header-nav-main__elements_link"
                        >
                            Найти друга
                        </Link>
                    </li>
                    <li className="header-nav-main__elements">
                        <a href="#" className="header-nav-main__elements_link">
                            Группы
                        </a>
                    </li>
                </ul>
            )}

            {userRegistered ? (
                <Link to={`/user/${userId}`} className="header-nav-sign">
                    Ваш пейдж
                </Link>
            ) : (
                <Link to="/sign" className="header-nav-sign">
                    Войти
                </Link>
            )}
        </div>
    );
};

export default HeaderNav;
