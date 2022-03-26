import React, { useState } from "react";
import "../sign-sass/user-profile-pop-up1.sass";

import { Link } from "react-router-dom";

const UserProfilePopUp1 = () => {
    return (
        <div className="user-profile-pop-up">
            <Link to="/sign/survey">
                <div className="user-profile-pop-up__container">
                    You have Registered!
                </div>
            </Link>
        </div>
    );
};

export default UserProfilePopUp1;