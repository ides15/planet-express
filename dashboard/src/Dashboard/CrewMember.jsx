import React from "react";

import { Row, Col, Statistic } from "antd";

const CrewMember = ({ member }) => {
    return (
        <Row gutter={24} className="member">
            <Col span={16}>
                <Statistic
                    title="Name"
                    value={`${member.first_name} ${member.last_name}`}
                    valueStyle={{ fontSize: "18px" }}
                />
            </Col>
            <Col>
                <Statistic
                    title="Age"
                    value={member.age}
                    valueStyle={{ fontSize: "18px" }}
                />
            </Col>
        </Row>
    );
};

export default CrewMember;
