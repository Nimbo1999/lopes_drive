import type { MetaFunction } from "@remix-run/node";
import { Outlet, useLocation } from "@remix-run/react";

import styles from "./layout.module.scss";
import { Anchor, Avatar } from "~/components";
import clsx from "clsx";

export const meta: MetaFunction = () => {
  return [
    { title: "Lopes Drive" },
    { name: "description", content: "Welcome to the lopes drive app!" },
  ];
};

export default function Index() {
  const location = useLocation();

  return (
    <div className={styles.layout}>
      <header className={styles.header}>
        <div className={styles.header__title}>
          <h1>Lopes Drive</h1>

          <Anchor
            to=".."
            className={clsx(styles["header__go-back"], {
              [styles["header__go-back--hidden"]]: location.pathname === "/",
            })}
          >
            Go back
          </Anchor>
        </div>

        <Anchor to="profile">
          <Avatar src="/assets/profile.jpeg" alt="profile" />
        </Anchor>
      </header>

      <Outlet />
    </div>
  );
}
