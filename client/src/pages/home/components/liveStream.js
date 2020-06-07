import React, { Fragment } from "react";

export const LiveStream = (props) => {
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
    border: "1px solid #1C2833",
    borderRadius: "5px",
  };

  return (
    <Fragment>
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
