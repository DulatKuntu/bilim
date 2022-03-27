import "../mentors-sass/mentors-card.sass";

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

const MentorsCard = ({ name, interests, bio }) => {
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
        <div className="mentors-card">
            <div className="mentors-card-upper">
                <div className="mentors-card-upper__img_container">
                    <img
                        src={randomImage}
                        alt=""
                        className="mentors-card-upper__img"
                    />
                </div>

                <div className="mentors-card-upper__text">
                    <div className="mentors-card-upper__text_main">{name}</div>
                    <div className="mentors-card-upper__text_specialization">
                        {interests}
                    </div>

                    <div className="mentors-card-description">{bio}</div>
                </div>
            </div>
        </div>
    );
};

export default MentorsCard;
