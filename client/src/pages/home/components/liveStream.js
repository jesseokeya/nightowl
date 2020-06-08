import React, { Fragment, useState } from "react";
import { Spin, Space } from "antd";

export const LiveStream = (props) => {
  const [spinner, setSpinner] = useState(true)

  const divStyle = {
    margin: "0 auto",
    maxWidth: "100vw",
    width: "100%",
    height: "100%",
  };

  const iframeStyle = {
    margin: "0 auto",
    maxWidth: "100vw",
    width: "100%",
    height: "32rem",
    border: "1px solid white",
    borderRadius: "5px",
  };

  setTimeout(() => setSpinner(false), 2000)

  return (
    <Fragment>
      {spinner && <div style={{ textAlign: "center" }}>
        <br />
        <Space size="large">
          <Spin size="large" />
        </Space>
        <br />
      </div>}
      <div className="wrap" style={divStyle} data-auto-scroll="true">
        <iframe
          className="iframe"
          title="nightOwl"
          style={iframeStyle}
          src={props.source}
        />
      </div>
    </Fragment>
  );
};
