import "../profession-sass/university-card.sass";

const UniversityCard = ({ name, descr, score }) => {
    return (
        <div className="university-card">
            <div className="university-card__text">
                <div className="university-card__text_main">{name}</div>
                <div className="university-card__text_specialization">
                    {score} баллов
                </div>
            </div>
            <div className="university-card-description">{descr}</div>
        </div>
    );
};

export default UniversityCard;
