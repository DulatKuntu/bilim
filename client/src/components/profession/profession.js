import "./profession-sass/profession.sass";

import ProfessionCard from "./profession-comp/profession-card";
import UniversityCard from "./profession-comp/university-card";

const Profession = () => {
    return (
        <div className="profession">
            <div className="profession-cards">
                <ProfessionCard
                    name="системный программист"
                    descr="Основная работа заключается в написании программ под определенные операционные системы. "
                />
                <ProfessionCard
                    name="программирование игр"
                    descr="Компьютерные игры сейчас на пике популярности, вне зависимости от того, для какой платформы они написаны – персональный компьютер или ноутбук, игровая приставка, планшет или обыкновенный смартфон."
                />
            </div>

            <div className="profession-cards">
                <ProfessionCard
                    name="отрасль WEB-программирования"
                    descr="Перспективно развивающееся направление, которое является своеобразным показателем развития интернет технологий. "
                />
                <ProfessionCard
                    name="тестировщики и работники техподдержки"
                    descr="В большинстве случаев работа для студентов в сфере IT начинается с этих профессий."
                />
            </div>

            <h2 className="profession-header">Рекомендаций:</h2>

            <UniversityCard
                name="Astana IT"
                score="112"
                descr="Универ в Астане"
            />

            <UniversityCard name="NU" score="120" descr="Универ в Астане" />

            <UniversityCard name="МУИТ" score="102" descr="Универ в Алмате" />
        </div>
    );
};

export default Profession;
