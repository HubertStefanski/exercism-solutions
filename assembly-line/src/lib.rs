// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
// #![allow(unused)]

static RATE: f64 = 221.0;

pub fn production_rate_per_hour(speed: u8) -> f64 {
    let base_rate = (speed as f64) * RATE;

    // if speed > 0 && speed <= 4 {
    //     return base_rate * 1.0;
    // }
    // if speed >= 5 && speed <= 8 {
    //     return base_rate * 0.90;
    // }
    // if speed >= 9 && speed <= 10 {
    //     return base_rate * 0.77;
    // };
    // base_rate

    match speed {
        5..=8 => base_rate * 0.9,
        9..=10 => base_rate * 0.77,
        _ => base_rate
    }
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    (production_rate_per_hour(speed) as u32) / 60
}
