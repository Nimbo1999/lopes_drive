import type { MetaFunction } from "@remix-run/node";
import { Outlet } from "@remix-run/react";

import styles from "./layout.module.scss";
import { Avatar } from "~/components";

export const meta: MetaFunction = () => {
  return [
    { title: "Lopes Drive" },
    { name: "description", content: "Welcome to the lopes drive app!" },
  ];
};

export default function Index() {
  return (
    <div className={styles.layout}>
      <header className={styles.header}>
        <h1>Lopes Drive</h1>

        <Avatar src="/assets/profile.jpeg" alt="profile" />
      </header>

      <Outlet />
    </div>
  );
}
