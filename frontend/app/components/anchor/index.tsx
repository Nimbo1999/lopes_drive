import { type LinkProps, Link } from "@remix-run/react";

import Styles from "./anchor.module.scss";
import clsx from "clsx";

export interface AnchorProps extends LinkProps {
  variant?: "link" | "button";
  color?: "primary" | "secondary";
  size?: "default" | "small";
}

export function Anchor({
  className,
  variant = "link",
  color = "primary",
  size = "default",
  ...props
}: AnchorProps) {
  return (
    <Link
      className={clsx(
        {
          [Styles.link]: variant === "link",
          [Styles.button]: variant === "button",
          [Styles.primary]: color === "primary",
          [Styles.secondary]: color === "secondary",
          [Styles.small]: size === "small",
        },
        className
      )}
      {...props}
    />
  );
}
