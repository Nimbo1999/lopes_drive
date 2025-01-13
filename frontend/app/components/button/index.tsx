import { ButtonHTMLAttributes } from "react";
import clsx from "clsx";

import styles from "./button.module.scss";

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "default" | "text" | "bordered";
  size?: "default" | "small";
  fullWidth?: boolean;
  color?: "primary";
}

export function Button({
  children,
  className,
  variant = "default",
  size = "default",
  fullWidth,
  color = "primary",
  ...props
}: Readonly<ButtonProps>) {
  return (
    <button
      {...props}
      className={clsx(className, styles.button, {
        [styles["button--small"]]: size === "small",
        [styles["button--full-width"]]: fullWidth,
        [styles["button--default"]]: variant === "default",
        [styles["button--text"]]: variant === "text",
        [styles["button--primary"]]: color === "primary",
      })}
    >
      {children}
    </button>
  );
}
