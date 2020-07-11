import React, { useState, useEffect } from "react";

import { Statistic, Card, Row, Col, Spin } from "antd";
import ApolloClient from "apollo-boost";
import { gql } from "apollo-boost";

import "./styles.scss";
import planetExpressShip from "./planet-express-ship.png";
import Delivery from "./Delivery";
import CrewMember from "./CrewMember";

const Dashboard = () => {
    const client = new ApolloClient({
        uri: "/graphql",
    });

    const [loading, setLoading] = useState(true);
    const [ship, setShip] = useState({});
    const [crew, setCrew] = useState({});
    const [deliveries, setDeliveries] = useState({});

    const getPlanetExpressData = async () => {
        const data = await client.query({
            query: gql`
                {
                    ship {
                        name
                        location
                        fuel_level
                    }
                    crew {
                        size
                        mood
                        captain_marooned
                        members {
                            id
                            first_name
                            last_name
                            age
                        }
                    }
                    allDeliveries {
                        id
                        timestamp
                        name
                        weight
                        fragile
                    }
                }
            `,
        });

        setShip(data.data.ship);
        setCrew(data.data.crew);
        setDeliveries(data.data.allDeliveries);
        setLoading(false);
    };

    useEffect(() => {
        getPlanetExpressData();
    }, []);

    if (loading) {
        return (
            <div className="dashboard">
                <div className="container">
                    <Spin size="large" />
                </div>
            </div>
        );
    }

    return (
        <div className="dashboard">
            <div className="container">
                <Card title="Ship" className="statistics">
                    <Row gutter={16}>
                        <Col span={8}>
                            <Statistic title="Name" value={ship.name} />
                        </Col>
                        <Col span={8}>
                            <Statistic title="Location" value={ship.location} />
                        </Col>
                        <Col span={8}>
                            <Statistic
                                title="Fuel Level"
                                value={ship.fuel_level}
                            />
                        </Col>
                    </Row>
                </Card>

                <img src={planetExpressShip} alt="Planet Express ship" />

                <Card title="Crew" className="statistics">
                    <Row gutter={16}>
                        <Col span={8}>
                            <Statistic title="Size" value={crew.size} />
                        </Col>
                        <Col span={8}>
                            <Statistic title="Mood" value={crew.mood} />
                        </Col>
                        <Col span={8}>
                            <Statistic
                                title="Captain Marooned?"
                                value={crew.captain_marooned}
                            />
                        </Col>
                    </Row>
                    <Row gutter={16} style={{ marginTop: "3rem" }}>
                        <Col span={24}>
                            {crew.members.map((member) => (
                                <CrewMember member={member} key={member.id} />
                            ))}
                        </Col>
                    </Row>
                </Card>

                <Card title="Deliveries" className="statistics">
                    <Row gutter={16}>
                        <Col span={24}>
                            {deliveries.map((delivery) => (
                                <Delivery
                                    delivery={delivery}
                                    key={delivery.id}
                                />
                            ))}
                        </Col>
                    </Row>
                </Card>
            </div>
        </div>
    );
};

export default Dashboard;
