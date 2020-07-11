import React from "react";

import { Row, Col, Statistic } from "antd";
import moment from "moment";

const Delivery = ({ delivery }) => {
    return (
        <Row gutter={24} className="member">
            <Col span={12}>
                <Statistic title="Name" value={delivery.name} />
            </Col>
            <Col span={6}>
                <Statistic
                    title="Weight"
                    value={delivery.weight}
                    suffix="lbs"
                />
            </Col>
            <Col span={6}>
                <Statistic
                    title="Timestamp"
                    value={moment(delivery.timestamp * 1000).format(
                        "MMM Do, YYYY @ hh:mma"
                    )}
                    valueStyle={{ fontSize: "16px" }}
                />
            </Col>
        </Row>
    );
};

export default Delivery;
