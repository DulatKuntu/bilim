import React, { useState } from "react";
import "../sign-sass/user-profile-pop-up.sass";

import { Link } from "react-router-dom";

const UserProfilePopUp = () => {
    return (
        <div className="user-profile-pop-up">
            <Link to="/">
                <div className="user-profile-pop-up__container">
                    You have Registered!
                </div>
            </Link>
        </div>
    );
};

export default UserProfilePopUp;
