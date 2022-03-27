import React, { useState } from "react";

import SignIn from "./sign-comp/sign-in";
import SignUp from "./sign-comp/sign-up";

import "./sign-sass/sign.sass";

import signMain from "./sign-img/main.png";

const Sign = ({ setUserRegistered, setUserId, user }) => {
    const [sign, setSign] = useState(false);

    return (
        <div className="sign">
            {sign ? (
                <SignUp setSign={setSign} user={user} />
            ) : (
                <SignIn
                    setUserRegistered={setUserRegistered}
                    setSign={setSign}
                    setUserId={setUserId}
                    user={user}
                />
            )}
        </div>
    );
};

export default Sign;
