import React, { useState } from "react";
import { Layout, Menu } from "antd";
import {
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  HomeOutlined,
  VideoCameraOutlined,
  SettingOutlined,
} from "@ant-design/icons";
import { Switch, Route, Link } from "react-router-dom";
import "../../styles/app.css";

import { LiveStream } from "./components";
import Resources from "../resources";

const { Header, Sider, Content } = Layout;

function Home() {
  const mappings = {
    "/": "1",
    "/resources": "2",
    "/settings": "3",
  };

  const pathname = window.location.pathname;
  const [collapsed, setCollapsed] = useState(true);
  const toggle = () => setCollapsed(!collapsed);

  return (
    <Layout>
      <Sider trigger={null} collapsible collapsed={collapsed}>
        <div
          className="logo"
          style={{
            textAlign: "center",
            marginTop: "8px",
            marginBottom: "5px",
          }}
        >
          <img
            src="/favicon.ico"
            alt="Smiley face"
            height="30"
            width="30"
            style={{ marginTop: "17px" }}
          />
        </div>
        <Menu
          theme="dark"
          mode="inline"
          defaultSelectedKeys={mappings[pathname]}
        >
          <Menu.Item key="1" icon={<HomeOutlined />}>
            <Link to="/">Home</Link>
          </Menu.Item>
          <Menu.Item key="2" icon={<VideoCameraOutlined />}>
            <Link to="/resources">Resources</Link>
          </Menu.Item>
          <Menu.Item key="3" icon={<SettingOutlined />}>
            <Link to="/settings">Settings</Link>
          </Menu.Item>
        </Menu>
      </Sider>
      <Layout>
        <Header className="site-layout-background" style={{ padding: 0 }}>
          {React.createElement(
            collapsed ? MenuUnfoldOutlined : MenuFoldOutlined,
            {
              className: "trigger",
              onClick: toggle,
              style: {
                marginLeft: "15px",
                fontWeight: "bold",
                fontSize: "25px",
                color: "white",
              },
            }
          )}
        </Header>
        <Content className="app-content">
          <div className="container">
            <Switch>
              <Route exact path="/">
                <LiveStream source="http://192.168.0.24:8081" />
              </Route>
              <Route path="/resources">
                <Resources />
              </Route>
              <Route path="/settings"></Route>
            </Switch>
          </div>
        </Content>
      </Layout>
    </Layout>
  );
}

export default Home;
