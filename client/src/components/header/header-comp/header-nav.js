import React from "react";

import "../header-sass/header-nav.sass";

import { Link } from "react-router-dom";

import logo from "../header-img/logo.png";

const HeaderNav = ({ userRegistered }) => {
    return (
        <div className="header-nav">
            <div className="header-nav-logo__container">
                <img src={logo} alt="" className="header-nav-logo" />
            </div>

            <div className="header-nav-company">EVISION</div>

            <ul className="header-nav-main">
                <li className="header-nav-main__elements">
                    <Link
                        to="/mentors"
                        className="header-nav-main__elements_link"
                    >
                        Менторы
                    </Link>
                </li>
                <li className="header-nav-main__elements">
                    <a href="#" className="header-nav-main__elements_link">
                        Стади Бадди
                    </a>
                </li>
                <li className="header-nav-main__elements">
                    <a href="#" className="header-nav-main__elements_link">
                        Группы
                    </a>
                </li>
            </ul>

            {userRegistered ? (
                <Link to="/mentors/1" className="header-nav-sign">
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
