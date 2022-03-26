import "../mentors-sass/mentors-card.sass";

const MentorsCard = () => {
    return (
        <div className="mentors-card">
            <div className="mentors-card-upper">
                <div className="mentors-card-upper__img_container">
                    <img src="" alt="" className="mentors-card-upper__img" />
                </div>

                <div className="mentors-card-upper__text">
                    <div className="mentors-card-upper__text_main">Test</div>
                    <div className="mentors-card-upper__text_specialization">
                        Test
                    </div>

                    <div className="mentors-card-description">
                        Lorem ipsum dolor sit amet consectetur adipisicing elit.
                        Accusamus, eum maiores ab
                    </div>
                </div>
            </div>
        </div>
    );
};

export default MentorsCard;
